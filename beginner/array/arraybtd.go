package main

import "fmt"

var productName [4]string
var price [4]float32

func main() {
	productName[0] = "MacBook"
	productName[1] = "iPad"

	price := [4]float32{30000.01, 20000.02}

	fmt.Println(productName)
	fmt.Println(price)
}
