package main

/*
Executable package type generates an executable file
Reusable packages are used as helpers (libraries, components, etc)

name of package determines package type:
- main = executable (must have a main function)
- any other name - reusable
*/

import "fmt"

/*
import libraries from other packages
fmt (format) = standard library package
by default main has access to no libraries unless imported
golang.org/pkg
*/

func main() {
	fmt.Println("Hi there!")
	printState()
}
