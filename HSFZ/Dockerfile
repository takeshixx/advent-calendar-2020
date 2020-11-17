FROM python:3

RUN apt-get update -y
RUN apt-get install python3 python3-pip python3-dev git libssl-dev libffi-dev build-essential -y
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
ENTRYPOINT python3 /server.py -c /config/config.toml -f $SECRET