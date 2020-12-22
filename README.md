# Advent Calendar of Advanced Cyber Fun

The 2020 edition with even more cyber fun. Wow!

## Intro

A CTF-like advent calendar that opens a port everyday, starting from port 1. The challenges incorporate different protocols and services ranging from ancient RFCs to bleeding edge technologies. Each port is meant to be solvable rather easily so that it doesn't take too much time.

The 2020 edition of the advent calendar was hosted at `xmas.rip`. The page contents are in the [_website](_website) directory.

## Prerequisites

Each task should run in a Docker container, similar to the previous [iteration](https://github.com/takeshixx/advent-calendar-2018). An exception are services that require to run on the host system because they are implemented in iptable rules or require specific Kernel features.

## Ideas

Tick the boxes to indicate the service has been implemented. Strikethrough text means challenge is already on the agenda.

- [x] ~~[RFC2965](https://tools.ietf.org/html/rfc2965): Implement webapp that requires Cookie2 HTTP header.~~
- [x] ~~[RFC7231](https://tools.ietf.org/html/rfc7231): Build a crawler with the From header.~~
- [x] ~~WebAssembly page with Golang ([xmas-webasm](xmas-webasm))~~
- [ ] Email stuff with DMARC/SPF and other stuff no one really understands.
- [x] ~~gRPC ([xmasgreetings](xmasgreetings))~~
- [ ] WebRTC (Servers send image, have to build the client)
- [ ] OData (e.g. with [godata](https://github.com/crestonbunch/godata))
- [ ] [Oauth flow](https://www.ory.sh/run-oauth2-server-open-source-api-security/)
- [ ] [ASN.1](https://golang.org/pkg/encoding/asn1/)
- [ ] [DNS over JSON](https://www.rfc-archive.org/getrfc?rfc=8427&tag=Representing-DNS-Messages-in-JSON)
- [ ] [Concise Binary Object Representation (CBOR)](https://tools.ietf.org/html/rfc7049) ([node-cbor](https://github.com/hildjj/node-cbor))
- [ ] GitHub Actions
  - Provide some code with a simple vulnerability, something that might be hard to compile (maybe in a weird programming language)
  - Users have to createa a GitHub Action that builds that code (without the vuln) and upload the binary as artifact
  - Allow users to submit a link to their build artifact
  - Create a own GitHub action that fetches the build artifact, runs it and checks if the vuln is still in there (restrict to GitHub URL)
  - In case binary works and the vuln is fixed, create a security issue in the repo with the "flag" (security issues should only be visible to repo owners)
- [x] ~~[DTLS 1.2](dtls/)~~
- [ ] TLS 1.3
- [x] ~~Webshop Race Conditions (Vouchers ausstellen) [WebRace](WebRace)~~
- [ ] [Diameter](http://www.freediameter.net/trac/)
- [ ] Websockets
- [x] ~~[HSZF(DoIP)/UDS](HSFZ/)~~
- [x] ~~[ELF](elf/) binary with wrong entry point. Prints the flag if entry point is corrected (see ELF folder)~~
- [x] ~~[PCAP_poly](PCAP_poly) PCAP File containg a Polyglot file containing the flag~~
- [x] ~~[xmas-socks](xmas-socks)~~
- [x] ~~[XMAS Cloud](xmas-cloud/), [Demo](http://svento-xmascloud.azurewebsites.net/)~~
- [x] ~~[proto](proto/): The proto challenge offers the token if you negotiate the correct [TLS ALPN protocol](https://en.wikipedia.org/wiki/Application-Layer_Protocol_Negotiation).~~
- [ ] [WireGuard](https://www.wireguard.com/)
- [x] ~~Server which requires to send specific UTF16 strings with correct BOM ([xmas-karaoke](xmas-karaoke))~~
- [x] ~~Image with Red Star OS watermark that includes the flag~~
- [x] ~~[XMAS scan](https://nmap.org/book/scan-methods-null-fin-xmas-scan.html) port that only allows packets with FIN, PSH and URG flags set~~
- [x] ~~[NTS](nts): [RFC 8915 - Network Time Security (NTS)](https://tools.ietf.org/html/rfc8915) server ([nts](nts))~~
- [x] ~~[SANTAS NAUGHTY LIST](./santas-naughty-list) is using a strict Content Security Policy to protected against all(?) XSS attacks.~~
- [x] ~~[HELLO XMAS/3.0](./http3only) is a HTTP/3.0 only server on UDP.~~
- [x] ~~[Santa's Christmas Factory](./santas-christmas-factory) is a web server affected by JavaScript prototype pollution.~~
- [x] ~~[Something with emojis](./emoji) is a small TCP server with a emoji "puzzle/quiz".~~

### Overall Challenge

The overall challenge will include an Amazon gift card again. Each port has a secret, the SHA256 hash of all secrets combined will be the password for an encrypted text on the website that includes instructions for receiving the gift card. Unfortunately we cannot just include the code of a gift card, because e.g. German gift cards won't work for Amazon Canada.

All keys are available at [_challenge/keys.json](_challenge/keys.json), the password in [_challenge/password](_challenge/password). Both have been generated with the [_challenge/generate_keys.py](_challenge/generate_keys.py) script.

### HealthState

Healthstate can be monitored with ``docker events --filter event=health_status``

## Agenda

First one or two ports should be fairly simple to give participants an easy start. The bold and underlined days are 2nd to 4th advents (1st is not in december this year) and they should have special challenges (harder/more complex).

| Port | Challenge | Path |
| ---- | --------- | ---- |
| 1    | A challenge that opens two web ports, port 1 returns a `Set-Cookie2` header with a port list that includes 11111. Send cookie to this port in `Cookie2` header according to [RFC2965](https://tools.ietf.org/html/rfc2965). | [xmas-cookie2](xmas-cookie2)
| 2    | A simple DTLv1.2 server that returns the secret. | [dtls](dtls)
| 3    | ELF binary with wrong entry point. If entry point is set to the correct one, the flag is printed. Binary is provided via web server. | [ELF](elf)
| 4    | xmas-socks is a simple portable parallel secure SOCKS server written in Go. | [xmas-socks](xmas-socks)
| 5    | A custom TLS server written in Go that returns the flag if you negotiate the correct [TLS ALPN protocol](https://en.wikipedia.org/wiki/Application-Layer_Protocol_Negotiation). | [proto](proto)
| <ins>**6**</ins>   | Web version of VSCode (Monaco editor) which allows to execute OS commands to read and decrypt flag via JavaScript. | [XMAS Cloud](xmas-cloud/)
| 7    | PCAP File containg a Polyglot file containing the flag a.k.a as Матрешка (Matreshka). | [PCAP_poly](PCAP_poly)
| 8    | High Speed Fahrzeugzugang (HSFZ) server where user's have to send a proper HSFZ packet that starts the car. | [HSFZ](HSFZ)
| 9    | A simple webserver that only shows the flag with a correct Request Context, which has to include a proper `From` and `Referer` header. | [xmas-from](xmas-from)
| 10    | A JPG file with a Red Star OS watermark that includes the flag. | [redstar](redstar)
| 11    | A karaoke service where clients have to reflect song lyrics in the UTF encoding indicated by the returned BOM. | [xmas-karaoke](xmas-karaoke)
| 12    | A simple [gRPC](https://grpc.io/) service where clients have to call the `XmasGreeting()` function with the `xmas` name. Protobuf definition will be provided. | [xmasgreetings](xmasgreetings)
| <ins>**13**</ins>    | WebAssembly page that requires a password. Prints the token with the proper password. | [xmas-webasm](xmas-webasm)
| 14    | A web shop with a race condition vulnerability. | [WebRace](WebRace)
| 15    | A Network Time Security service which returns the token in a NTPv4 Server Negotiation Message.| [nts](nts)
| 16    | A simple IP-HTTPS server where a client has to bring up a IP-HTTPS link. | [ip-https](ip-https)
| 17    | A HTTP server that is only accessible via [TLS-over-SCTP](https://tools.ietf.org/html/rfc3436). | [tls-over-sctp](tls-over-sctp)
| 18    | A [Rust](https://www.rust-lang.org/) RPC service. | [rustRPC](rustRPC)
| 19    | A HTTP/3-only server on UDP. | [http3only](http3only)
| <ins>**20**</ins>    | CSP bypass challenge. | [santas-naughty-list](santas-naughty-list)
| 21    | Simple FTP server with login and a secret file. | [xmas-ftpd](xmas-ftpd)
| 22    | JavaScript type pollution challenge. | [Santa's Christmas Factory](./santas-christmas-factory)
| 23    | TCP server with emoji puzzles/quizzes. | [Something with emojis](./emoji)
| 24    | XMAS scan port that returns the token in a ICMP 13 packet. Can be solved with Nmap XMAS scan and Wireshark. | [xmas-tcpflags](xmas-tcpflags)
