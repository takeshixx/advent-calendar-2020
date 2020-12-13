package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "Welcome to XMAS Edge Services\n")
		return
	}
	if len(r.Header["Content-Length"]) > 0 && r.Header["Content-Length"][0] != "" {
		if r.Header["Content-Length"][0] == "18446744073709551615" {
			fmt.Fprintf(w, "IP-HTTPS link brought up........ (%s)\n", os.Getenv("XMAS_SECRET"))
		} else {
			fmt.Fprintf(w, "Best edge services. Good prices, very much SLA. Wow.\n")
		}
	} else {
		http.Error(w, "Invalid content-length, with this you will get nowhere.\n", http.StatusForbidden)
	}
}
