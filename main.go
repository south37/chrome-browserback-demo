package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/index", handleIndex)
	http.HandleFunc("/other", handleOther)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	acceptMediaType := r.Header["Accept"]
	if len(acceptMediaType) == 1 && acceptMediaType[0] == "application/json" {
		handleIndexJSON(w, r)
	} else {
		handleIndexHTML(w, r)
	}
}

func handleIndexJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"message":"Hello from an index (application/json)"}`)
}

func handleIndexHTML(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>Index Page</title>
    <script src="/index.js"></script>
  <body>
    <div>Hello from an Index Page!</div>
	<a href="/other">Other Page</a>
  </body>
</html>
`)
}

func handleOther(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>Other Page</title>
  <body>
    <div>Hello from an Other Page!</div>
  </body>
</html>
`)
}
