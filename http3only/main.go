package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"sync"

	_ "net/http/pprof"

	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)

var flagToken *string

type binds []string

func (b binds) String() string {
	return strings.Join(b, ",")
}

func (b *binds) Set(v string) error {
	*b = strings.Split(v, ",")
	return nil
}

func defaultPageHandler(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Flag string
	}{
		*flagToken,
	}
	t, e := template.ParseFiles("templates/index.html")
	if e != nil {
		log.Printf("Error parsing template:%v\n", e)
		return
	}
	e = t.Execute(w, data)
	if e != nil {
		log.Printf("Error executing template:%v\n", e)
	}
}

func setupHandler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(defaultPageHandler))

	fsStatic := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fsStatic))

	return mux
}

func main() {
	bs := binds{}
	flag.Var(&bs, "bind", "bind to")
	flagToken = flag.String("flag", "PLEASE PASS THE FLAG via --flag", "flag")
	tcp := flag.Bool("tcp", false, "also listen on TCP")
	flag.Parse()

	if len(bs) == 0 {
		bs = binds{"localhost:6121"}
	}

	handler := setupHandler()
	quicConf := &quic.Config{}

	var wg sync.WaitGroup
	wg.Add(len(bs))
	for _, b := range bs {
		bCap := b
		go func() {
			var err error
			if *tcp {
				log.Printf("listen on UDP/TCP %s", bCap)
				certFile, keyFile := generateTLSConfig()
				err = http3.ListenAndServe(bCap, certFile, keyFile, handler)
			} else {
				log.Printf("listen on UDP %s", bCap)
				server := http3.Server{
					Server:     &http.Server{Handler: handler, Addr: bCap},
					QuicConfig: quicConf,
				}
				err = server.ListenAndServeTLS(generateTLSConfig())
			}
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// Setup a bare-bones TLS config for the server
func generateTLSConfig() (string, string) {
	return "cert/fullchain0.pem", "cert/privkey0.pem"
}
