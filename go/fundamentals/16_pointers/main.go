// memory
// function call in go are "pass by value"
// a copy of each arguments is made regardless of size
// potentially slow for large data structure
// more difficult to manage program state

// pointers
// variables that "point to" memory
// value of the variable itself is a memory address
// accessing data requires dereferencing the pointer

package main

import "fmt"

func incrementPointer(x *int) {
	// dereferencing or inderecting
	*x += 1
}

func incrementValue(x int) int {
	return x + 1
}

// case: slice
func appendArrayValue(x []int, v int) {
	x = append(x, v)
}

func appendArrayPointer(x *[]int, v int) {
	*x = append(*x, v)
}

func updateArrayValue(x []int, i int, v int) {
	x[i] = v
}

func updateArrayPointer(x *[]int, i int, v int) {
	(*x)[i] = v
}

func main() {
	value := 10
	// asterisk (*) when used with type indicates value is a pointer
	var pointerValue *int
	// ampersand (&) create a pointer from variable
	pointerValue = &value

	fmt.Println(pointerValue)
	fmt.Println(*pointerValue)

	i := 1

	// create variable that has a memory address that points to variable (will not have a copy)
	incrementPointer(&i)
	fmt.Println(i)

	i = incrementValue(i)
	fmt.Println(i)

	slice := []int{1, 2, 3}

	// appendArrayPointer(&slice, 5)
	appendArrayValue(slice, 4)
	fmt.Println(slice)

	// updateArrayPointer(&slice, 0, 5)
	// updateArrayValue(slice, 0, 4)
	// fmt.Println(slice)
}
