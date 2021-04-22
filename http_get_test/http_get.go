package httpget

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", login)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	r.ParseForm()
	username := r.Form["username"]
	pwd := r.Form["password"]
	fmt.Fprintf(w, username[0])
	fmt.Fprintf(w, "<br>")
	fmt.Fprintf(w, pwd[0])
}
