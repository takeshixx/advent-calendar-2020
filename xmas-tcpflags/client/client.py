#!/usr/bin/env python3
import sys
import socket
import random

from scapy.all import sr1, TCP, IP


if __name__ == '__main__':
    if len(sys.argv) < 2:
        print('Usage: {} host port interface'.format(sys.argv[0]))
        sys.exit(1)
    host = sys.argv[1]
    port = int(sys.argv[2])
    interface = sys.argv[3]
    # RSTs to this port must be blocked
    #sport = random.randint(1024,65535)
    sport = 49113
    ip = IP(dst=host)
    xmas = TCP(sport=sport, dport=port, flags='FPU', seq=1000)
    resp = sr1(ip / xmas, iface=interface)
    if resp:
        print(resp.show2())
    else:
        print('no response')
