package main

import "os"

type Page struct {
	Title string
	Body  []byte // abyte slice
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func main() {

}
