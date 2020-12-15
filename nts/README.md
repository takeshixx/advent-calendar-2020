# NTS

A [RFC 8915 - Network Time Security (NTS)](https://tools.ietf.org/html/rfc8915) server. NMap does not detect the NTS service, yet. The token is disclosed in the [NTPv4 Server Negotiation](https://tools.ietf.org/html/rfc8915#section-4.1.7) message.

## Description

Santa needs to deliver its presents on time. However, there are bad people out there trying to convince Santa that it's not Christmas time, yet. Attempts to [secure it](https://tools.ietf.org/html/rfc5906) in the past [had failed](https://www.semanticscholar.org/paper/Analysis-of-the-NTP-Autokey-Procedures-R%C3%B6ttger/a1781712cec129d5c7311a915e4d0076117ee33f). Now, Santa hopefully found a better solution to get the correct time.

## Build with Docker

```bash
docker-compose build
```

## Run with Docker

The flag can be found in the patch [xmas-patch.diff](xmas-patch.diff). Edit the ``xmas-patch.diff`` and modify the string ``"the-token-is.TOKEN.xmas.rip"``. Then rebuild the container:

```bash
docker-compose up --build
```

The NTS server is exposed on TCP port 4460. You can change it in the ``docker-compose.yml`` file. Please note that the NTP port (123/UDP) does not need to be exposed, to 

## Solution

Setup a client which speaks the NTS protocol and find [NTPv4 Server Negotiation](https://tools.ietf.org/html/rfc8915#section-4.1.7) message in the server response:

A solution for *ntpsec* can be found in this repo: 

    docker run -it --entrypoint /usr/local/sbin/ntpd nts_nts  -c /etc/ntp.d/config_client -n

You might need to set the correct host IP in the line ``server 172.17.0.1 nts noval`` in the [config_client](config_client) and rebuild the container.

It will output:

    2020-12-15T08:51:09 ntpd[1]: NTSc: Using server the-token-is.TOKEN.xmas.rip=>51.195.44.86

## References

- https://blog.cloudflare.com/nts-is-now-rfc/
- https://gitlab.com/NTPsec/ntpsec/-/releases
- https://docs.ntpsec.org/latest/NTS-QuickStart.html
- Buildfile based on: https://github.com/AevaOnline/docker-ntpsec