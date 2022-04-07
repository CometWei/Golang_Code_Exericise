package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/", teapot)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	r.ParseForm()
	username := r.Form["user"]
	pwd := r.Form["pass"]
	fmt.Fprintf(w, username[0])
	fmt.Fprintf(w, "<br>")
	fmt.Fprintf(w, pwd[0])
}

func teapot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	r.ParseForm()
	w.WriteHeader(http.StatusTeapot)
	w.Write([]byte("I'm a Teapot!!"))
}
