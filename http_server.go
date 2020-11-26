package main

import (
	"net/http"
)

func main() {

	r := http.NewServeMux()
	r.HandleFunc("/", handleRequest)
	r.HandleFunc("/formResult", handleResponse)
	http.ListenAndServe(":9090", r)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html")

	htmlForm := `
		<form action="/formResult" method="POST">
			<input name="first_name" />
			<input type="submit" value="OK" />
		</form>
	`
	// returnedStream := html.EscapeString(r.URL.Query().Get("key"))
	println(htmlForm)
	_, err := w.Write([]byte(htmlForm))
	if err != nil {
		println(err)
	}
}

func handleResponse(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic("error")
	}

	firstName := r.PostForm.Get("first_name")
	w.Write([]byte(firstName))
}
