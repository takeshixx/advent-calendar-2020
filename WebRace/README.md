# Santa's Shop



## Running

Insert flag in Dockerfile.

Maybe due to Django we need to adjust the settings for Production with Allowed_hosts in ``WebRace/webapp/source/app/conf/production/settings.py`` and insert our DNS name

## Solution

```bash
echo -e '\x00\x00\x00\x04\x00\x01\x13\x37\x24\x12'| nc 127.0.0.1 6801
```
## Building & Running

```bash
make day14-build
sudo docker-compose up day14
```