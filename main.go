// every go application is structured into packages
// so every go file that you're going to have
// is going to declare what package it's a part of
// main is a special package,
// which is the entry point of your application
// 一個package會包含一個或多個.go結束的源代碼文件。
package main

import (
	"fmt" // packages allows us format string
)

// entry point of our application
// main function in the main package
func main() {
	name := "Go Developers"
	fmt.Println("Hello for", name)
}
