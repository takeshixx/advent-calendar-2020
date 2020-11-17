import pwn


def pack_all(data):
    p = pwn.make_packer('all', endian='big')
    return p(data)

def unpack_all(data):
    return pwn.unpack(data, len(data)*8, endian='big', sign=True)
