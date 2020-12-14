# NTS

[RFC 8915 - Network Time Security (NTS)](https://tools.ietf.org/html/rfc8915) server. NMap does not detect NTS service, yet.

## Description

Santa needs to deliver its presents on time. However, there are bad people out there trying to convince Santa that it's not Christmas time, yet. Attempts to [secure it](https://tools.ietf.org/html/rfc5906) in the past [had failed](https://www.semanticscholar.org/paper/Analysis-of-the-NTP-Autokey-Procedures-R%C3%B6ttger/a1781712cec129d5c7311a915e4d0076117ee33f). Now, Santa hopefully found a better solution to get the correct time.

## Build with Docker

```bash
docker-compose build
```

## Run with Docker

The flag can be found in the TLS certifcate. Edit the ``openssl.cnf`` and add the line ``OU=EXAMPLE_TOKEN``. Then run ``./generate-cert.sh`` to generate the certificate.

Run the container:

```bash
docker-compose up --build
```

The NTS server is exposed on TCP and UDP port 4460. You can change it in the ``docker-compose.yml`` file.

## Solution

    TODO

## References

- https://gitlab.com/NTPsec/ntpsec/-/releases
- https://docs.ntpsec.org/latest/NTS-QuickStart.html
- Buildfile based on: https://github.com/AevaOnline/docker-ntpsec
- https://blog.cloudflare.com/nts-is-now-rfc/