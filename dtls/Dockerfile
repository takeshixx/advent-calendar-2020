FROM golang:1.14

RUN apt-get update -y
RUN apt-get install git libssl-dev build-essential -y

COPY ./advent.patch /

RUN git clone https://github.com/pion/dtls.git

RUN cd dtls && git apply /advent.patch && go build -o /server examples/listen/selfsign/main.go

ENV SECRET replacemeplease

EXPOSE 4444/udp
ENTRYPOINT /server
#ENTRYPOINT python2.7 /server.py