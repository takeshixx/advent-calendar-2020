FROM python:3

RUN apt-get update -y
RUN apt-get install python3 python3-pip python3-dev git libssl-dev libffi-dev build-essential netcat -y
RUN python3 -m pip install --upgrade pip
RUN python3 -m pip install --upgrade pwntools toml

ENV SECRET replacemeplease

ENV PWNLIB_NOTERM=true

COPY ./server.py /
RUN mkdir /config
RUN mkdir /helper
COPY ./config/config.toml /config/config.toml
COPY ./config/load_config.py /config/load_config.py
COPY ./helper/helper.py /helper/helper.py
COPY ./helper/Packets.py /helper/Packets.py

HEALTHCHECK --interval=900s --timeout=10s --retries=3 CMD echo -e '\x00\x00\x00\x04\x00\x01\x13\x37\x24\x12'| nc 127.0.0.1 6801 -q 1 | grep 'BRUM BRUM' || exit 1

ENTRYPOINT python3 /server.py -c /config/config.toml -f $SECRET