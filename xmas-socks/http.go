package main

import (
	"fmt"
	"log"
	"net/http"
)

func serveHTTP(dir string, port int) {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(dir))
	mux.Handle("/", fs)

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: mux,
	}
	log.Printf("Listening on :%d...\n", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
