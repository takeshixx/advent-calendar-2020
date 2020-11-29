# Website

The main advent website.

## Building

```bash
sudo docker build --tag advent-website .
```

## Running

```bash
sudo docker run --restart=always -v /etc/letsencrypt:/usr/local/etc/letsencrypt advent-website
```