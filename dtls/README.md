# Santa secure UDP

TODO fancy text for DTLS

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
