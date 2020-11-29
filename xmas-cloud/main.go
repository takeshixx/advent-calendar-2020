package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
)

func execute(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	code := string(body)
	runCode(code, w)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	mime.AddExtensionType(".css", "text/css; charset=utf-8")
	mime.AddExtensionType(".js", "application/javascript; charset=utf-8")
	mime.AddExtensionType(".html", "text/html")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fsMonaco := http.FileServer(http.Dir("./node_modules/monaco-editor"))
	http.Handle("/monaco-editor/", http.StripPrefix("/monaco-editor/", fsMonaco))

	http.HandleFunc("/api/exec", execute)
	http.HandleFunc("/api/hello", hello)
	http.HandleFunc("/api/headers", headers)

	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
