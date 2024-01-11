package main

import "fmt"

func main() {
	var i int = 8 // i is int = 8
	fmt.Println("initial value of i is", i)
	fmt.Println("address of i", &i)
	var pointer_integer *int
	pointer_integer = &i
	fmt.Println("address of pointer is", pointer_integer)
	*pointer_integer = 5
	fmt.Println("new value of i is", i)
}
