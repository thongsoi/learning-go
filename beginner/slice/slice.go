package main

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4, 5} //slice declaration using []type{....}
	fmt.Println(slc[0])
	fmt.Println(slc[1])
	fmt.Println(slc[2])
	fmt.Println(slc[3])
	fmt.Println(slc[4])
	fmt.Println(len(slc), cap(slc)) // slice length, slice capacity

}
