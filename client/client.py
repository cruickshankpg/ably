import time

import grpc
import stateless_pb2
import stateless_pb2_grpc
import stateful_pb2
import stateful_pb2_grpc
import uuid
import hashlib


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
                print("conndecting to: {}".format(self.server_address))
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
        for gen in stub.GenerateSequence(req):
            yield gen


def main():
    # total = 0
    # client = StatelessClient()
    # for num in client.get_sequence(15):
    #     print("received: " + str(num))
    #     total += num
    #
    # print("total: " + str(total))
    client = StatefulClient(port=9000)
    seq = ""
    recv_check = b''
    for item in client.get_sequence(5):
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


if __name__ == '__main__':
    main()
