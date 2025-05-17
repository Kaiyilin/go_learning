package main

import (
	"fmt"
)

func main() {
	// syntax
	// slice_name := []datatype{values}

	my_empty_slices := []int{}
	my_slices := []int{1, 2, 3}

	fmt.Println("my_empty_slices:", my_empty_slices)
	fmt.Println("my_empty_slices capacity:", cap(my_slices))
	fmt.Println("my_empty_slices len:", len(my_slices))

	fmt.Println("my_slices:", my_slices)
	fmt.Println("my_slices capacity:", cap(my_slices))
	fmt.Println("my_slices len:", len(my_slices))

	// I can also create slices from arrays
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	my_slice_from_arr := arr1[2:4]

	fmt.Printf("my_slice_from_arr: %v\n", my_slice_from_arr)               // [12 13]
	fmt.Printf("my_slice_from_arr capacity: %d\n", cap(my_slice_from_arr)) // 4
	fmt.Printf("my_slice_from_arr len: %d\n", len(my_slice_from_arr))      // 2
	// note that the length of the slice is 2, but the capacity is 4
	// because the slice is created from an array of size 6, and the slice starts at index 2
	// so the slice can grow to the end of the array, which is 4 elements

	// slice_name := make([]type, length, capacity)

	my_make_slice := make([]int, 5, 10)
	fmt.Printf("my_make_slice: %v\n", my_make_slice)
	fmt.Printf("my_make_slice len: %d\n", len(my_make_slice))
	fmt.Printf("my_make_slice capacity: %d\n", cap(my_make_slice))

	// with omitted capacity, capacity is equal to length
	my_make_slice_2 := make([]int, 5)
	fmt.Printf("my_make_slice_2: %v\n", my_make_slice_2)
	fmt.Printf("my_make_slice_2 capacity: %d\n", cap(my_make_slice_2))
	fmt.Printf("my_make_slice_2 len: %d\n", len(my_make_slice_2))
}
