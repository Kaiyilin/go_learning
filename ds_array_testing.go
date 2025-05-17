package main

import (
	"fmt"
)

// array got fixed size
// go is 0-based index
func main() {
	var cars_branding = [5]string{"Volvo", "BMW", "Ford", "Mazda", "Honda"}

	fmt.Print(cars_branding, "\n")
	fmt.Println("first element of the array:", cars_branding[0]) // print the first element of the array

	// switch the last element to Nissan
	fmt.Println("Array before changes:", cars_branding)
	cars_branding[4] = "Nissan"
	fmt.Println("Array after changes:", cars_branding)
}
