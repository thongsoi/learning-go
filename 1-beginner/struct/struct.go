package main

import "fmt"

// declar struct Circle with 3 fields
type Circle struct {
	X, Y, R int
}

func main() {
	c1 := Circle{2, 3, 5}
	c2 := Circle{Y: 1, R: 5}
	c3 := Circle{}

	fmt.Println(c1)
	fmt.Println(c1.X)
	fmt.Println(c2)
	fmt.Println(c2.R)
	fmt.Println(c3)

}
