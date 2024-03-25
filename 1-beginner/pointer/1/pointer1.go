package main

import "log"

func main() {
	myString := "Green"
	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("After func is called, myString is set to", myString)
}

func changeUsingPointer(s *string) {

	newValue := "Red"
	*s = newValue
}
