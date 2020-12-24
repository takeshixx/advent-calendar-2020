# Rust RPC

Note:
call Class impl ``XmasServer`` with function ``xmas`` with the ``secret`` string parameter ``XMAS``


## Running

Insert flag in Dockerfile.

```bash
docker build -t dayxx_rustrpc .
sudo docker run -d -p xx:8080 dayxx_rustrpc
```

## Solution

Run the XmasClient

```bash
./XmasClient --server_addr xmas.rip:xx --secret "XMAS"
```

## Hint

Newest version of https://github.com/google/tarpc/
