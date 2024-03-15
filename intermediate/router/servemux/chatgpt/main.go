package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	var tmplt = template.Must(template.ParseFiles("templates/index.html"))
	if err := tmplt.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Template execution error:", err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	var tmplt = template.Must(template.ParseFiles("templates/login.html"))
	if err := tmplt.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Template execution error:", err)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	var tmplt = template.Must(template.ParseFiles("templates/register.html"))
	if err := tmplt.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Template execution error:", err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Listening...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
