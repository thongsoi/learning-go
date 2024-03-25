package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Java", "Python"}
	fmt.Println(courseName)
	courseName = append(courseName, "C", "C#", "HTML", "CSS", "JS") //add C to slice with append
	fmt.Println(courseName)
	courseWeb := courseName[4:7] //select only last 3 arrays
	fmt.Println(courseWeb)
	courseWeb = courseName[0:3] //select only first 3 arrays
	fmt.Println(courseWeb)
}
