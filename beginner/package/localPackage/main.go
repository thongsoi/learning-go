package main

import (
	"fmt"
	wow "localPackage/myPackages" //moduel localPackage/myPackage folder
)

func main() {
	fmt.Println("Package in Go") //from standard package fmt.Println function
	wow.Wow1()                   //from local package wow.wow1 function
	wow.Wow2()                   //from local package wow.wow2 function
}
