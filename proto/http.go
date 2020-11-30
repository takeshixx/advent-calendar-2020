package main

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/net/http2"
)

var (
	certBytes   []byte
	cert        *x509.Certificate
	key         *ecdsa.PrivateKey
	h2TLSConfig *tls.Config
	certPool    = x509.NewCertPool()
	flag        string
)

// runServer runs a TLS connection that indicates support for HTTP/2 and for
// arbitrary protocol "foo" negotiated via TLS ALPN. The server contains a
// mapping of this protocol name to handleFoo. For demonstrating HTTP request
// handling, it implements a route that simply returns "gotcha" back to the
// client.
func runServer() {

	h2TLSConfig = &tls.Config{
		Certificates: []tls.Certificate{
			tls.Certificate{
				Certificate: [][]byte{certBytes},
				PrivateKey:  key,
			},
		},
		RootCAs:    certPool,
		ServerName: tlsCommonName,
		ClientCAs:  certPool,
		NextProtos: []string{
			"h2",
			"xmas",
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var versionString string
		switch r.TLS.Version {
		case tls.VersionTLS10:
			versionString = "1.0"
		case tls.VersionTLS11:
			versionString = "1.1"
		case tls.VersionTLS12:
			versionString = "1.2"
		case tls.VersionTLS13:
			versionString = "1.3"
		}

		if r.TLS.NegotiatedProtocol != "" {
			w.Write([]byte(fmt.Sprintf("You've connected with TLS version %s and negotiated the protocol: %s\n", versionString, r.TLS.NegotiatedProtocol)))
		} else {
			w.Write([]byte(fmt.Sprintf("You've connected with TLS version %s and negotiated no protocol.\n", versionString)))
		}
	})

	ln, err := net.Listen("tcp", listenAddress)
	if err != nil {
		fmt.Printf("error starting listener: %v\n", err)
		os.Exit(1)
	}
	tlsLn := tls.NewListener(ln, h2TLSConfig)

	server := &http.Server{
		Addr:    listenAddress,
		Handler: mux,
		TLSNextProto: map[string]func(*http.Server, *tls.Conn, http.Handler){
			"xmas": handleXMASProto,
		},
	}

	http2.ConfigureServer(server, nil)

	log.Printf("listen on %v\n", listenAddress)
	err = server.Serve(tlsLn)
	log.Fatalf("error starting server: %v\n", err)
}

func handleXMASProto(server *http.Server, conn *tls.Conn, handler http.Handler) {
	log.Printf("xmas proto connection from %s\n", conn.RemoteAddr().String())

	wtr := bufio.NewWriter(conn)
	_, err := wtr.WriteString(fmt.Sprintf("You've connected via xmas protocol. You found the flag: %s\n", flag))
	if err != nil {
		log.Printf("error writing flag: %v\n", err)
	}
	wtr.Flush()
	if err != nil {
		log.Printf("error writing flag: %v\n", err)
	}
	conn.Close()
}

func generateCert() {
	var err error
	key, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("error generating key: %v\n", err)
		os.Exit(1)
	}

	template := &x509.Certificate{
		Subject: pkix.Name{
			CommonName: tlsCommonName,
		},
		DNSNames: []string{tlsCommonName},
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
			x509.ExtKeyUsageClientAuth,
		},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageKeyAgreement | x509.KeyUsageCertSign,
		SerialNumber:          big.NewInt(1),
		NotBefore:             time.Now().Add(-1 * time.Second),
		NotAfter:              time.Now().Add(1 * time.Hour * 24 * 365),
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	certBytes, err = x509.CreateCertificate(rand.Reader, template, template, key.Public(), key)
	if err != nil {
		log.Fatalf("error generating self-signed cert: %v\n", err)
	}

	cert, err = x509.ParseCertificate(certBytes)
	if err != nil {
		log.Fatalf("error parsing generated certificate: %v\n", err)
	}

	certPool.AddCert(cert)
}

func readFlag() {
	flagB, err := ioutil.ReadFile(flagPath)
	if err != nil {
		log.Fatalf("couldn't load flag from %s: %v", flagPath, err)
	}
	flag = string(flagB)
}

func startServer(cmd *cobra.Command, args []string) {
	readFlag()
	generateCert()
	runServer()
}
