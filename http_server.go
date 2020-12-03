package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func main() {

	r := http.NewServeMux()
	r.HandleFunc("/", handleRequest)
	r.HandleFunc("/formResult", handleResponse)
	r.HandleFunc("/auth", authBasic)
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

func authBasic(w http.ResponseWriter, r *http.Request) {
	authBase64 := r.Header.Get("Authorization")
	tempSlice := strings.Split(string(authBase64), " ")
	if len(tempSlice) != 2 {
		w.WriteHeader(400)
		return
	}

	encoding := base64.StdEncoding
	decodedAuth, err := encoding.DecodeString(string(tempSlice[1]))
	if err != nil {
		fmt.Errorf("%v", err)
		w.WriteHeader(400)
		return
	}

	authString := strings.Split(string(decodedAuth), ":")
	if len(authString) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if authString[0] != "aladdin" && authString[1] != "opensesame" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte("OK"))
}
