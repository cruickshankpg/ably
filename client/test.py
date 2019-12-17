import hashlib

import client
import subprocess
import os


def test_stateless():
    pid = subprocess.Popen(["../stateless-server", "-debug"], stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL).pid

    stateless = client.StatelessClient(port=9001)
    count = 0
    for val in stateless.get_sequence(2):
        count += 1

    os.kill(pid, 0xf)
    if count != 2:
        print("stateless test failed, too many terms")
    else:
        print("stateless test passed")


def test_stateful():
    pid = subprocess.Popen(["../stateful-server", "-debug"], stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL).pid

    stateful = client.StatefulClient(port=9000)
    seq = ""
    recv_check = b''
    for item in stateful.get_sequence(4):
        seq = seq + str(item.number)
        if len(item.checksum) is not None:
            recv_check = item.checksum
    sha = hashlib.sha256(bytes(seq, 'utf8'))
    csum = sha.digest()

    os.kill(pid, 0xf)

    if csum != recv_check:
        print("stateful test failed, checksums are different")
        print("received: {}".format(recv_check))
        print("calculated: {}".format(csum))
    else:
        print("stateful test passed")


if __name__ == '__main__':
    test_stateless()
    test_stateful()
