package main

import (
	"fmt" // packages allows us format string
)

var a int
var b int = 2
var c = 3

// := cannot be used outside of function
// e := 5 // syntax error: non-declaration statement outside function body
// e is not declared in the package scope

func main() {
	var student_name string // declare first
	student_name = "myName" // assign later
	fmt.Println(student_name)
	fmt.Println(a) // 0
	fmt.Println(b) // 2
	fmt.Println(c) // 3
	a = 1
	fmt.Println(a) // 1
	d := 4         // short declaration
	fmt.Println(d) // 4
	// d = "test" // cannot use "test" (untyped string constant) as int value in assignment
}
