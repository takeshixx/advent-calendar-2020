import struct
import pwn
from enum import Enum


class Identifier(Enum):
    REQUEST = 1
    RESPONSE = 2
    WTF_IDENTIFIER = 0x40
    WTF_IDENTIFIER2 = 0x43


# Common Response Codes
class Crc(Enum):
    SERVICE_NOT_SUPPORTED = 0x7F


class Sessions(Enum):
    DEFAULT_SESSION = 1
    PROGRAMMING_SESSION = 2
    EXTENDED_DIAGNOSTIC_SESSION = 3
    DEVELOPMENT_SESSION = 4


class ErrorCodes(Enum):
    GR = 0x10       # General Reject
    SNS = 0x11      # Service not supported
    SFNS = 0x12     # Sub-Function not supported
    IMLOIF = 0x13   # Incorrect message length or invalid format
    BRR = 0x21      # Busy repeat request
    CNC = 0x22      # Conditions not correct
    RSE = 0x24      # Request sequence error
    ROOR = 0x31     # Request out of range
    RCRRP = 0x78    # Request correctly received, but response is pending
    SFNSIAS = 0x7E  # Sub-Function not supported in active session
    SNSIAS = 0x7F   # Service not supported in active session


class UDSPacket:

    sid_without_sub = [0x14, 0x23, 0x24, 0x2A, 0x2E, 0x2F, 0x34, 0x35, 0x36, 0x37, 0x3D, 0x84]

    def __init__(self):
        self.data = None
        self.__crc = None
        self.__error_code = None
        self.__sid = None
        self.__sub = None
        self.__parameter = None

    def parse(self, packet):
        return self._parse_uds(packet)

    def create(self, packet):
        if not len(packet):
            return self
        self.data = packet
        if len(packet) > 1:
            self._parse_uds(packet)
        else:
            self.__sid = packet[0]
        return self

    @property
    def crc(self):
        return self.__crc

    @crc.setter
    def crc(self, crc):
        self.__crc = crc

    @property
    def sid(self):
        return self.__sid

    @sid.setter
    def sid(self, sid):
        self.__sid = sid

    @property
    def sub(self):
        return self.__sub

    @sub.setter
    def sub(self, sub):
        self.__sub = sub

    def get_full_packet(self):
        if len(self.data) == 0:
            return b''
        if self.crc is ErrorCodes.SNSIAS.value:
            return self.data
        ret = pwn.p8(self.sid)
        if self.sub is not None:
            ret += pwn.p8(self.sub)
        if self.parameter is not None:
            ret += self.parameter
        return ret

    @property
    def error_code(self):
        return self.__error_code

    @error_code.setter
    def error_code(self, error_code):
        self.__error_code = error_code

    @property
    def parameter(self):
        return self.__parameter

    @parameter.setter
    def parameter(self, parameter):
        self.__parameter = parameter

    # Parsing UDS from Standard of 2016 UDS
    def _parse_uds(self, packet):
        self.data = packet
        if len(packet) == 0:
            self.__parameter = None
            self.__sub = None
            self.__crc = None
            self.__sid = None
            return self
        if packet[0] is ErrorCodes.SNSIAS.value:
            self.__crc = packet[0]
        else:
            self.__sid = packet[0]
            if len(packet) > 1:
                if self.__sid in self.sid_without_sub:
                    self.__parameter = packet[1:]
                else:
                    if packet[1] == 0x41:
                        #print("WARNING: The SID: {} is not in the list of no subfunctions. "
                        #      "However it occured that the sub function is the value 0x41. "
                        #      "Thus assuming its a parameter".format(self.__sid))
                        self.__parameter = packet[1:]
                        return self
                    self.__sub = packet[1]
                    self.__parameter = packet[2:]
        return self


class HSZFPacket:

    def __init__(self):
        self.length = None
        self.identifier = None
        self.uds_packet = None
        self.header = None
        self.reciever = None
        self.sender = None
        self.session = None

    def parse(self, packet):
        if len(packet) < 7:
            return
        self.header = packet[:8]
        length, identifier, sender, reciever = self._parse_hszf(self.header)
        self.length = length
        self.sender = sender
        self.reciever = reciever
        self.identifier = Identifier(identifier)
        #print(packet[8:6+self.length])
        self.uds_packet = UDSPacket().parse(packet[8:6+self.length])
        return self

    def create(self, payload, sender, reciever):
        self.uds_packet = UDSPacket().create(payload)
        self.sender = sender
        self.reciever = reciever
        self.header = self._create_hszf_header(payload, sender, reciever)
        return self

    @staticmethod
    def _parse_hszf(response):
        return struct.unpack(">IHBB", response)

    @staticmethod
    def _create_hszf_header(payload, sender, reciever):
        return pwn.p32(len(payload)+2, endian='big') + pwn.p16(1, endian='big') + pwn.p8(sender) + pwn.p8(reciever)

    def get_uds_error_code(self):
        return self.uds_packet.error_code

    def get_uds_error(self):
        return self.uds_packet.crc

    @staticmethod
    def break_rna(seq_rna, *break_point):
        """
        Splits the string at the given offsets into a list
        :param seq_rna:
        :param break_point:
        :return:
        """
        seq_rna_list = []
        no_of_breakpoints = len(break_point)
        for _ in range(no_of_breakpoints):
            for index in break_point:
                seq_rna_list.append(seq_rna[:index])
                seq_rna = seq_rna[index:]
            break
        seq_rna_list.append(seq_rna)
        return seq_rna_list

    def get_readable_packet(self):
        splitted_packet = self.break_rna((self.header + self.uds_packet.get_full_packet()).hex(), 8, 4, 2, 2)
        return ' '.join(splitted_packet)

    def get_full_packet(self):
        return self.header + self.uds_packet.get_full_packet()
