# Request Context: From

A simple webserver that only shows the flag with a correct Request Context, which has to include a proper `From` and `Referer` header.

## Description

```html
Santa only comes here after visiting the /christmas_market from his own xmas.rip domain. His elfes therefore always track the proper <a href="https://tools.ietf.org/html/rfc7231#section-5.5">Request Context</a>.
```

## Solution

```bash
curl -H "Referer: /christmas_market" -H "From: santa@xmas.rip" xmas.rip:123
```

## Building

```bash
sudo docker build --tag advent-xmas-from .
```

## Running

```bash
sudo docker run -p 8000:8000 advent-xmas-from
```