package main

import (
	"fmt"
	wowow "package/mypackages"
)

func main() {
	fmt.Println("Package in Go")

	// now you use the name of the package specified by the "package" keyword inside of those files from *amazing_package* folder
	wowow.PrintWowow()
	wowow.PrintAnotherWow()
	fmt.Println("and this is Some_var from the same imported package:", wowow.Some_var)
}
