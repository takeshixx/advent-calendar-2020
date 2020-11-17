import pwn

import argparse
from config.load_config import load_config
from helper.Packets import HSZFPacket, Identifier
from helper.helper import *
from time import sleep


class Server:

    def __init__(self, config, secret):
        self.config = config
        self.secret = secret
        self.wait_for_response = config['gateway']['waitForResponse']
        s = pwn.server(bindaddr=config['gateway']['host'], port=config['gateway']['port'], callback=self.callb)
        self.server_conn = s.next_connection()

    @staticmethod
    def no_unpack(data):
        return data

    @staticmethod
    def not_implemented(data, r, pk):
        return b'\x7F'+pack_all(data)+b'\x11'

    @staticmethod
    def emulate_device_busy(data, r, pk):
        response = HSZFPacket()
        response.create(b'\x7F'+pack_all(data)+b'\x78', pk.sender, pk.reciever)
        r.send(response.get_full_packet())
        pwn.log.info("Sending 'Request correctly received, but response is pending'")
        sleep(2)
        return b'\xCC\xFF\x12\xFF'

    def sid_switch(self, case):
        return {
            
            0x0105: lambda data, r, pk: b'\x7F'+pack_all(data)+b'\x13',
            0x010541: lambda data, r, pk: b'\x59\x02\xFF\x02\xFF\x65\x50\x48\x38\x0C\x50\x48\x38\x0D\x50\x48\x38',
            0x08: lambda data, r, pk: b'\x7F'+self.switch_pack(bytes_needed(data))(data)+b'\x21',
            0x02: lambda data, r, pk: b'\x7F'+pack_all(data)+b'\x13',
            0x07: self.emulate_device_busy,
            0x00: lambda data, r, pk: b'WELCOME',
            0x01: lambda data, r, pk: b'WELCOME',
            0x12: lambda data, r, pk: b'YOU_ARE_ONTO_SOMETHING',
            0x1224: lambda data, r, pk: b'WROOM WROOM ' + bytes(self.secret, encoding='utf8'),
            0x24: lambda data, r, pk: b'YOU_ARE_ONTO_SOMETHING',
            0x2412: lambda data, r, pk: b'BRUM BRUM ' + bytes(self.secret, encoding='utf8'),
            0x14: lambda data, r, pk: b'\x7F'+pack_all(data)+b'\x13',
            0x33: lambda data, r, pk: b'\x7F'+pack_all(data)+b'\x10', # General Reject if it doesnt match ECU ident
        }.get(case, self.not_implemented)

    def callb(self, r):
        while True:
            client_input = r.recvrepeat(self.wait_for_response)
            if len(client_input) < 1:
                sleep(0.1)
                continue
            pwn.info("Recieved: " + pwn.hexdump(client_input))
            pk = HSZFPacket()
            pk.parse(client_input)
            if pk.identifier == Identifier.REQUEST:
                uds_data = pk.uds_packet.get_full_packet()
                if len(uds_data) < 1:
                    data = 0x00
                else:
                    data = unpack_all(uds_data)
                # ECU identifier
                if pk.reciever != 0x37 or pk.sender != 0x13:
                    response_data = self.sid_switch(0x33)(0x33, r, pk)
                else:
                    response_data = self.sid_switch(data)(data, r, pk)
                response = HSZFPacket()
                response.create(response_data, pk.reciever, pk.sender)
                pwn.info("Sending the response: " + pwn.hexdump(response.get_full_packet()))
                r.send(response.get_full_packet())

                if data == 0x22F100:
                    pwn.info("--Got Session Request--\n" + pwn.hexdump(pk.uds_packet.get_full_packet()))
                    response = HSZFPacket()
                    response.create(b'\x62\xF1\x00\x01\x81\x00\x01', pk.reciever, pk.sender)
                    pwn.info("Sending the response: " + pwn.hexdump(response.get_full_packet()))
                    r.send(response.get_full_packet())

            r.clean_and_log()


def main(arg):
    config = load_config(arg.config)
    Server(config, arg.flag)


if __name__ == "__main__":
    # execute only if run as a script
    parser = argparse.ArgumentParser()
    parser.add_argument("--config", "-c", help="TOML Configuration File")
    parser.add_argument("--flag", "-f", help="the secret flag")
    args = parser.parse_args()
    main(args)



