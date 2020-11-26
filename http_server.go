package main

import "net/http"

func main() {

	r := http.NewServeMux()
	r.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", r)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	println(r.RemoteAddr)
	returnedStream := "Hello World!"
	_, err := w.Write([]byte(returnedStream))
	if err != nil {
		println(err)
	}
}
