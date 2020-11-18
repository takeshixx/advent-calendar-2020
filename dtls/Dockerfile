FROM golang:1.14

RUN apt-get update -y
RUN apt-get install git libssl-dev build-essential -y

COPY ./advent.patch /

RUN git clone https://github.com/pion/dtls.git

RUN cd dtls && git apply /advent.patch && go build -o /server examples/listen/selfsign/main.go

ENV SECRET replacemeplease

EXPOSE 4444/udp

HEALTHCHECK --interval=900s --timeout=10s --retries=3 CMD echo "Q" | openssl s_client -dtls1_2 -connect 127.0.0.1:4444 2>/dev/null | grep 'Secret:' || exit 1

ENTRYPOINT /server
