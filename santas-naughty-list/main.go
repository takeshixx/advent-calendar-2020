package main

import (
	"fmt"
	"html/template"
	"log"
	"mime"
	"net/http"
)

const port = 8080

func cspMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'none'; style-src 'self'; script-src 'self';")
		next.ServeHTTP(w, r)
	})
}

func langHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Lang string
	}{
		r.URL.Query().Get("lang"),
	}
	t, e := template.ParseFiles("templates/lang.js")
	if e != nil {
		log.Printf("Error parsing template:%v\n", e)
		return
	}
	e = t.Execute(w, data)
	if e != nil {
		log.Printf("Error executing template:%v\n", e)
	}
}

func defaultPageHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Lang    string
		Naughty string
		// Comment template.HTML
		DataAction template.HTMLAttr
	}{
		r.URL.Query().Get("lang"),
		r.URL.Query().Get("naughty"),
		// template.HTML("<!-- TODO: <script src=\"/lang.js?lang=\"></script> -->"),
		template.HTMLAttr("data-action='{\"method\":\"deleteEntry\",\"parameters\":\"{{index}}\"}'"),
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

func main() {
	mime.AddExtensionType(".css", "text/css; charset=utf-8")
	mime.AddExtensionType(".js", "application/javascript; charset=utf-8")
	mime.AddExtensionType(".html", "text/html")

	mux := http.NewServeMux()

	mux.Handle("/lang.js", cspMiddleware(http.HandlerFunc(langHandler)))
	mux.Handle("/default.aspx", cspMiddleware(http.HandlerFunc(defaultPageHandler)))
	mux.Handle("/", cspMiddleware(http.RedirectHandler("/default.aspx?naughty=%5B%0A%20%20%20%20%7B%20%22firstName%22%3A%20%22COVID%22%2C%20%22lastName%22%3A%20%2219%22%20%7D%2C%0A%20%20%20%20%7B%20%22firstName%22%3A%20%22Donald%22%2C%20%22lastName%22%3A%20%22Trump%22%20%7D%0A%5D&lang=en", http.StatusFound)))

	fsStatic := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", cspMiddleware(http.StripPrefix("/static/", fsStatic)))

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
