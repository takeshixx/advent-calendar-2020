package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	if len(r.Header["Referer"]) > 0 && r.Header["Referer"][0] != "" &&
		len(r.Header["From"]) > 0 && r.Header["From"][0] != "" {
		if strings.ToLower(r.Header["Referer"][0]) == "/christmas_market" && strings.ToLower(r.Header["From"][0]) == "santa@xmas.rip" {
			fmt.Fprintf(w, "Welcome Santa, here is your secret: %s\n", os.Getenv("XMAS_SECRET"))
		} else {
			fmt.Fprintf(w, "Nice try bugger, maybe next time!!\n")
		}
	} else {
		http.Error(w, "You are not Santa. Where are you even From?! Go away!!\n", http.StatusForbidden)
	}
}
