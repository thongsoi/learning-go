package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Golang App %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	log.Println("Starting http server on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("error starting http server:", err)
		return
	}
}
