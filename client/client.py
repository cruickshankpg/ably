"""client.py connects to a specified stateful or stateless number generation
server and verifies the results given.

Usage:
    client.py [-l LENGTH] stateful <port>
    client.py [-l LENGTH] stateless <port>

Options:
    -l LENGTH  sequence length. If not specified a random number will be chosen between 0 and 0xffff.
"""

import time

import docopt
import grpc
import stateless_pb2
import stateless_pb2_grpc
import stateful_pb2
import stateful_pb2_grpc
import uuid
import hashlib
import random


class AblyClient(object):
    def __init__(self, host='localhost', port=9001, max_timeout=16):
        self.server_address = "{}:{}".format(host, port)
        self.max_timeout = max_timeout

    def get_sequence(self, sequence_len, timeout=1):
        raise NotImplemented


class StatelessClient(AblyClient):
    def get_sequence(self, sequence_len, timeout=1):
        last = 0
        while True:
            try:
                channel = grpc.insecure_channel(self.server_address)
                stub = stateless_pb2_grpc.StatelessNumberGeneratorStub(channel)

                req = stateless_pb2.GenerateSequenceRequest(restartFrom=last)

                for gen in stub.GenerateSequence(req):
                    sequence_len -= 1
                    last = gen.number
                    yield last
                    if sequence_len == 0:
                        channel.close()
                        return
            except grpc.RpcError as e:
                print("hit exception: " + e.__str__())
                if timeout > self.max_timeout:
                    raise e
                time.sleep(timeout)
                timeout *= 2


class StatefulClient(AblyClient):
    def __init__(self, port=9000, host='localhost', max_timeout=16):
        super().__init__(port=port, host=host, max_timeout=max_timeout)

    def get_sequence(self, sequence_len, timeout=1):
        conn_id = uuid.uuid4()
        count = 0
        while True:
            try:
                channel = grpc.insecure_channel(self.server_address)
                stub = stateful_pb2_grpc.StatefulNumberGeneratorStub(channel)

                if count == 0:
                    stream = self.__first_gen_iter(stub, conn_id, sequence_len)
                else:
                    stream = self.__recon_gen_iter(stub, conn_id, count - 1)

                for gen in stream:
                    count += 1
                    yield gen
                    if gen.finalItem:
                        channel.close()
                        return
            except grpc.RpcError as e:
                if e.code() == grpc.StatusCode.NOT_FOUND:
                    raise e
                print("hit exception: " + e.__str__())
                if timeout > self.max_timeout:
                    raise e
                time.sleep(timeout)
                timeout *= 2

    @staticmethod
    def __first_gen_iter(stub, conn_id, sequence_len):
        req = stateful_pb2.GenerateSequenceRequest(connectionID=str(conn_id), sequenceLength=sequence_len)
        for gen in stub.GenerateSequence(req):
            yield gen

    @staticmethod
    def __recon_gen_iter(stub, conn_id, last_recv):
        req = stateful_pb2.ReconnectSequenceRequest(connectionID=str(conn_id), lastReceivedIndex=last_recv)
        for gen in stub.ReconnectSequence(req):
            yield gen


def run_stateful(port, sequence_len):
    client = StatefulClient(port=port)
    seq = ""
    recv_check = b''
    for item in client.get_sequence(sequence_len):
        seq = seq + str(item.number)
        if len(item.checksum) is not None:
            recv_check = item.checksum
    sha = hashlib.sha256(bytes(seq, 'utf8'))
    csum = sha.digest()
    if csum != recv_check:
        print("checksums are different")
        print("received: {}".format(recv_check))
        print("calculated: {}".format(csum))
    else:
        print("checksums match")


def run_stateless(port, sequence_len):
    total = 0
    client = StatelessClient(port=port)
    for num in client.get_sequence(sequence_len):
        total += num
    print("total: " + str(total))


def main():
    args = docopt.docopt(__doc__)

    random.seed()
    length = int(random.uniform(0, 0xffff))
    if args['-l'] is not None:
        length = int(args['-l'])

    if args['stateless']:
        print('starting stateless client with sequence length: {}'.format(length))
        run_stateless(args['<port>'], length)
    elif args['stateful']:
        print('starting stateful client with sequence length: {}'.format(length))
        run_stateful(args['<port>'], length)


if __name__ == '__main__':
    main()
