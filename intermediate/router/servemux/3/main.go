package main

import (
	"html/template"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	var tmplt = template.Must(template.ParseFiles("templates/login.html"))
	tmplt.Execute(w, nil)
}
func register(w http.ResponseWriter, r *http.Request) {
	var tmplt = template.Must(template.ParseFiles("templates/register.html"))
	tmplt.Execute(w, nil)
}

func main() {

	http.HandleFunc("/", login)
	http.HandleFunc("/register", register)

	log.Println("Listening...")
	http.ListenAndServe(":8000", nil)
}
