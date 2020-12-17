# TLS over SCTP

A simple HTTP server that is accessible via [TLS-over-SCTP](https://tools.ietf.org/html/rfc3436).

## Solution

```
ncat -vvlkp 9999 --sh-exec "ncat --ssl --sctp xmas.rip 17"
curl localhost:9999/token
```