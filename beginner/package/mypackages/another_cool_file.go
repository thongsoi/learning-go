// another_cool_file.go

package wowow

import "fmt"

var Some_var = 42 // this will also be automatically exported

func PrintAnotherWow() {
	fmt.Println("this is also from wowow package, but from a different file")
}
