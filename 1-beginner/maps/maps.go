package main

import "fmt"

var product = make(map[string]float64) //[key]value

func main() {
	product["Macbook"] = 40000
	fmt.Println("product=", product)

}
