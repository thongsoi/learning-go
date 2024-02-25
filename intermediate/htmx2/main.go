// main.go
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/login.html")
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is POST
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Get the username and password from the form data
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "admin" && password == "password" {
			w.Write([]byte("Logged in"))
		} else {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}
	})

	http.ListenAndServe(":8080", nil)
}
