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
- [x] [RFC7231](https://tools.ietf.org/html/rfc7231): Build a crawler with the From header. ([xnas-from](xmas-from))
- [ ] WebAssembly page with Golang
- [ ] Email stuff with DMARC/SPF and other stuff no one really understands.
- [ ] gRPC
- [ ] WebRTC (Servers send image, have to build the client)
- [ ] OData (e.g. with [godata](https://github.com/crestonbunch/godata))
- [ ] Oauth flow https://www.ory.sh/run-oauth2-server-open-source-api-security/
- [ ] ASN.1 (https://golang.org/pkg/encoding/asn1/)
- [ ] DNS over JSON (https://www.rfc-archive.org/getrfc?rfc=8427&tag=Representing-DNS-Messages-in-JSON)
- [ ] Concise Binary Object Representation (CBOR) https://tools.ietf.org/html/rfc7049 https://github.com/hildjj/node-cbor
- [ ] GitHub Actions
  - Provide some code with a simple vulnerability, something that might be hard to compile (maybe in a weird programming language)
  - Users have to createa a GitHub Action that builds that code (without the vuln) and upload the binary as artifact
  - Allow users to submit a link to their build artifact
  - Create a own GitHub action that fetches the build artifact, runs it and checks if the vuln is still in there (restrict to GitHub URL)
  - In case binary works and the vuln is fixed, create a security issue in the repo with the "flag" (security issues should only be visible to repo owners)
- [x] [DTLS 1.2](dtls/)
- [ ] TLS 1.3
- [ ] Webshop Race Conditions (Vouchers ausstellen)
- [ ] Diameter Protokoll (http://www.freediameter.net/trac/)
- [ ] Websockets
- [x] [HSZF(DoIP)/UDS](HSFZ/)
- [x] [ELF](elf/) binary with wrong entry point. Prints the flag if entry point is corrected (see ELF folder)
- [ ] PCAP File containg a Polyglot file containing the flag a.k.a as Матрешка (Matreshka)
- [x] [xmas-socks](xmas-socks)
- [x] [XMAS Cloud](xmas-cloud/), [Demo](http://svento-xmascloud.azurewebsites.net/)
- [x] [proto](proto/): The proto challenge offers the token if you negotiate the correct [TLS ALPN protocol](https://en.wikipedia.org/wiki/Application-Layer_Protocol_Negotiation).
- [ ] [WireGuard](https://www.wireguard.com/)

### Overall Challenge

We should have an overall challenge where users can win something. E.g. the first one to solve everything is able to decrypt a secret, which could be a gift card like last time. We shouldn't use RSA keys to encrypt the secret this time...

### HealthState

Healthstate can be monitored with ``docker events --filter event=health_status``

## Agenda

First one or two ports should be fairly simple to give participants an easy start. The bold days are 2nd to 4th advents (1st is not in december this year) and they should have special challenges (harder/more complex).

| Port | Challenge | Path |
| ---- | --------- | ---- |
| 1    | A challenge that opens two web ports, port 1 returns a `Set-Cookie2` header with a port list that includes 11111. Send cookie to this port in `Cookie2` header according to [RFC2965](https://tools.ietf.org/html/rfc2965). | [xmas-cookie2](xmas-cookie2)
| 2    | A simple DTLv1.2 server that returns the secret. | [dtls](dtls)
| 3    | ELF binary with wrong entry point. If entry point is set to the correct one, the flag is printed. Binary is provided via web server. | [ELF](elf)
| 4    | | [xmas-socks](xmas-socks)
| 5    | | [proto](proto)
| **6**   | | [XMAS Cloud](xmas-cloud/)
| 7    | | []()
| 8    | | [HSFZ](HSFZ)
| 9    | | []()
| 10    | | []()
| 11    | | []()
| 12    | | []()
| **13**    | | []()
| 14    | | []()
| 15    | | []()
| 16    | | []()
| 17    | | []()
| 18    | | []()
| 19    | | []()
| **20**    | | []()
| 21    | | []()
| 22    | | []()
| 23    | | []()
| 24    | | []()
