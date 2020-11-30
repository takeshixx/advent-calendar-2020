# Website

The main advent website.

## Building

```bash
sudo docker build --tag advent-website .
```

## Running

```bash
sudo docker run -d --restart=always -v /etc/letsencrypt:/etc/letsencrypt -p 80:80 -p 443:443 --name website advent-website
```