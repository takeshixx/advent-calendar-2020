package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/armon/go-socks5"
)

func startSOCKS() {
	// Create a SOCKS5 server
	conf := &socks5.Config{
		Rewriter: addressRewriter{},
	}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf("%s:%d", address, port)
	log.Printf("Starting SOCKS5 proxy on %s\n", addr)
	// Create SOCKS5 proxy
	if err := server.ListenAndServe("tcp", addr); err != nil {
		panic(err)
	}
}

type addressRewriter struct {
}

func (addressRewriter) Rewrite(ctx context.Context, request *socks5.Request) (context.Context, *socks5.AddrSpec) {
	log.Printf("Connection from %v\n", request.RemoteAddr.IP)
	if request.RemoteAddr.IP.IsLoopback() {
		return ctx, &socks5.AddrSpec{IP: net.IPv4(127, 0, 0, 1), Port: httpAuthorizedPort}
	}
	return ctx, &socks5.AddrSpec{IP: net.IPv4(127, 0, 0, 1), Port: httpAnonymousPort}
}
