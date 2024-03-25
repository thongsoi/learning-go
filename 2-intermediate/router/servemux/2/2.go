package main

import (
	"fmt"
	"log"
	"net/http"
)

type textHandler struct {
	responseText string
}

func (th *textHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, th.responseText)
}

type indexHandler struct {
}

func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"text/html",
	)
	html :=
		`<!DOCTYPE html>
		<html lang="en">
			<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>BiomassX</title>
			</head>
			<body>
		          <a href="/markets">Markets</a> |  <a href="/products">Products</a> |  <a href="/services">Services</a> 
        	</body>
		</html>`
	fmt.Fprint(w, html)

}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &indexHandler{})

	markets := &textHandler{"market1...n"}
	mux.Handle("/markets", markets)

	products := &textHandler{"product1...n"}
	mux.Handle("/products", products)

	services := &textHandler{"services1...n"}
	mux.Handle("/services", services)

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
