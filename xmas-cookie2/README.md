# XMAS Cookie2

A webserver that runs on two ports, supporting obsolete `Set-Cookie2` and `Cookie2` headers as defined in [RFC2965](https://tools.ietf.org/html/rfc2965). The first port returns a `Set-Cookie2` header with a port list that also includes a high-range port. The high range port is only accessible with the provided cookie in a `Cookie2` header.


## Description

```
Everybody loves christmas cookies! The best ones are freshly backed, but the old ones may do as well. Even if there SHOULD NOT be used anymore because their <a href="https://tools.ietf.org/html/rfc2965">receipes</a> are deprecated!
```

## Solution

Go to port 1 to get the `Set-Cookie2` header with the secret and then send it to port 11111:

```
curl -H "Cookie2: supasdersecretstuffhere" localhost:9999
```

# Building & Running

```
sudo docker build --tag day01 .
```

Run a test instance:

```
sudo docker run -it -p 8888:1 -p 9999:11111 day01
```