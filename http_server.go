package main

import (
	"html"
	"net/http"
)

func main() {

	r := http.NewServeMux()
	r.HandleFunc("/", handleRequest)
	http.ListenAndServe(":9090", r)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html")
	returnedStream := html.EscapeString(r.URL.Query().Get("key"))
	println(r.URL.Query().Get("key"))
	_, err := w.Write([]byte(returnedStream))
	if err != nil {
		println(err)
	}
}
