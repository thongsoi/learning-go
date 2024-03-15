package main

import (
	"net/http"
)

func main() {
	func index(w http.ResponseWriter, r *http.Request) {
		t := template.ParseFiles ("tmpl.html")
		t.Execute (w, "Hello World!")
	}
}

func main()  {
server := http.Server {
	Addr: "127.0.0.1:8080",

}
http.HandleFunc("/", index)
server.ListenAndServe()
}