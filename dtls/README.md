# Santa secure UDP

## Description

```html
Routing of Christmas presents requires secure transport, from the north pole to your living room. On a long journey, Santa has to do a lot of handshakes with several third parties to establish new connections. But because of COVID, Santa has to avoid handshakes all the way to stay healthy. Otherwise, Santa won't be able to sniff who made the best Christmas cookies.

Santa and helpers still have to provide secure services, so they started using a secure, yet unknown for most protocol on today's service they said is similar to TLS...?
```

## Running

Insert flag in Dockerfile.

```bash
docker build -t dayxx_dtls .
sudo docker run -d -p xx:4444/udp dayxx_dtls
```

## Solution

```bash
openssl s_client -dtls1_2 -connect <ip>:<port>
```
