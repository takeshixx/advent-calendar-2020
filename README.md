# Advent Calendar of Advanced Cyber Fun

Ideas for the 2020 edition

## Intro

A CTF-like advent calendar that opens a port everyday, starting from port 1. The challenges incorporate different protocols and services ranging from ancient RFCs to bleeding edge technologies. Each port is meant to be solvable rather easily so that it doesn't take too much time.

## Prerequisites

Each task should run in a Docker container, similar to the previous [iteration](https://github.com/takeshixx/advent-calendar-2018). An exception are services that require to run on the host system because they are implemented in iptable rules or require specific Kernel features.

## Ideas

- RFC2965: Implement webapp that requires Cookie2 HTTP header.
- RFC7231: Build a crawler with the From header.
- WebAssembly page with Golang
- Email stuff with DMARC/SPF and other stuff no one really understands.
- gRPC
- OData (e.g. with [godata](https://github.com/crestonbunch/godata))
- Oauth flow https://www.ory.sh/run-oauth2-server-open-source-api-security/
- ASN.1 (https://golang.org/pkg/encoding/asn1/)
- DNS over JSON (https://www.rfc-archive.org/getrfc?rfc=8427&tag=Representing-DNS-Messages-in-JSON)
- Concise Binary Object Representation (CBOR) https://tools.ietf.org/html/rfc7049 https://github.com/hildjj/node-cbor
- GitHub Actions
  - Provide some code with a simple vulnerability, something that might be hard to compile (maybe in a weird programming language)
  - Users have to createa a GitHub Action that builds that code (without the vuln) and upload the binary as artifact
  - Allow users to submit a link to their build artifact
  - Create a own GitHub action that fetches the build artifact, runs it and checks if the vuln is still in there (restrict to GitHub URL)
  - In case binary works and the vuln is fixed, create a security issue in the repo with the "flag" (security issues should only be visible to repo owners)
- DTLS 1.2
- Webshop Race Conditions (Vouchers ausstellen)
- Diameter Protokoll

### Overall Challenge

We should have an overall challenge where users can win something. E.g. the first one to solve everything is able to decrypt a secret, which could be a gift card like last time. We shouldn't use RSA keys to encrypt the secret this time...

## Agenda

| Port | Challenge | Path |
| ---- | --------- | ---- |
| 1    | | []()
| 2    | | []()
| 3    | | []()
| 4    | | []()
| 5    | | []()
| 6    | | []()
| 7    | | []()
| 8    | | []()
| 9    | | []()
| 10    | | []()
| 11    | | []()
| 12    | | []()
| 13    | | []()
| 14    | | []()
| 15    | | []()
| 16    | | []()
| 17    | | []()
| 18    | | []()
| 19    | | []()
| 20    | | []()
| 21    | | []()
| 22    | | []()
| 23    | | []()
| 24    | | []()
