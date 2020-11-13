# Santa's Little ELF

ELF binary with wrong entry point. If entry point is set to the correct one, the flag is printed. Binary is provided via web server. Insert flag in Dockerfile.

## Running

```bash
docker build -t dayxx_SLE .
sudo docker run -dit -p xx:80 dayxx_SLE
```
