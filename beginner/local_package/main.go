package main

import (
	"fmt"
	wow "local_package/mypackages" //moduel local_package/mypackage folder
)

func main() {
	fmt.Println("Package in Go") //from standard package fmt.Println function
	wow.Wow1()                   //from local package wow.wow1 function
	wow.Wow2()                   //from local package wow.wow2 function
}
