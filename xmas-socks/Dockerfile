FROM golang:1.15.5
RUN apt update && apt install -y curl proxychains4 dnsutils
RUN curl --output /proxy.sh https://raw.githubusercontent.com/takeshixx/tools/master/net/proxy.sh
RUN chmod +x /proxy.sh
WORKDIR /app
COPY anonymous ./anonymous/
COPY permitted ./permitted/
COPY . .
RUN echo "JAHr7BBNh0fp3IlbC2DItfWOV0CYIj3J" > /app/permitted/flag.txt
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o socks .
RUN chmod -R go+rX .
EXPOSE 9000
CMD ["./socks"]