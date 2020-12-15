package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

var defaultResp = `HTTP/1.1 200 HELO
Server: XMAS-EDGE-ROUTER-v24.12
Content-Length: 133
Content-Type: text/plain

Welcome to the XMAS Edge Services page. This is only for authorized users to remotely access internal resources. Please go away. Now.
`

var successResp = `HTTP/1.1 200 OK
Server: XMAS-EDGE-ROUTER-v24.12
Content-Length: 45
Content-Type: text/plain

You made it!`

var badRequestResp = `HTTP/1.1 400 NONO
Server: XMAS-EDGE-ROUTER-v24.12
Content-Length: 75
Content-Type: text/plain

What was that?! Please don't do that again. This is only for remote access.
`

var notFoundResp = `HTTP/1.1 404 DUNNO
Server: XMAS-EDGE-ROUTER-v24.12
Content-Length: 43
Content-Type: text/plain

Dunno, probably lost that page... ¯\_(ツ)_/¯
`

var methodNotAllowedResp = `HTTP/1.1 405 WHAT
Server: XMAS-EDGE-ROUTER-v24.12
Content-Length: 34
Content-Type: text/plain

What do you mean? Invalid method!!
`

func handleConnection(conn net.Conn) {
	defer conn.Close()
	req := make([]byte, 4096)
	for {
		n, err := conn.Read(req)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
				return
			}
		}

		cl := false
		if bytes.Contains(req, []byte("Content-Length: 18446744073709551615")) {
			cl = true
			req = bytes.ReplaceAll(req, []byte("Content-Length: 18446744073709551615"), []byte(""))
		}

		reqBuf := bytes.NewReader(req)
		reqReader := bufio.NewReader(reqBuf)
		pReq, err := http.ReadRequest(reqReader)
		if err != nil {
			n, err = conn.Write([]byte(badRequestResp))
			if err != nil {
				log.Println(n, err)
			}
			return
		}

		var resp string

		if pReq.URL.Path == "/" {
			if pReq.Method == "GET" {
				resp = defaultResp
			} else if pReq.Method == "POST" {
				resp = methodNotAllowedResp
			} else {
				resp = methodNotAllowedResp
			}
		} else if pReq.URL.Path == "/IPHTTPS" {
			if pReq.Method == "GET" {
				resp = methodNotAllowedResp
			} else if pReq.Method == "POST" {
				if cl {
					resp = successResp + " " + os.Getenv("XMAS_SECRET")
				} else {
					resp = badRequestResp
				}
			} else {
				resp = methodNotAllowedResp
			}
		} else {
			resp = notFoundResp
		}

		n, err = conn.Write([]byte(resp))
		if err != nil {
			return
		}
	}
}

func main() {
	log.SetFlags(log.Lshortfile)
	certPtr := flag.String("cert", "/etc/letsencrypt/archive/xmas.rip/fullchain1.pem", "certificate")
	keyPtr := flag.String("key", "/etc/letsencrypt/archive/xmas.rip/privkey1.pem", "key")
	listenPtr := flag.String("host", "0.0.0.0", "listening host")
	portPtr := flag.Int("port", 8000, "listening port")
	flag.Parse()

	cer, err := tls.LoadX509KeyPair(*certPtr, *keyPtr)
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", fmt.Sprintf("%s:%d", *listenPtr, *portPtr), config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
