// store multiple pieces of the same kind of data
// access item in the array using index (starts at 0)
// arrays is fixed size and cannot be resized

package main

import "fmt"

func main() {
	// creating an array
	var intArr1 [3]int
	intArr1[0] = 1
	intArr1[1] = 2
	intArr1[2] = 3
	fmt.Println(intArr1)

	intArr2 := [3]int{1, 2, 3}
	fmt.Println(intArr2)

	intArr3 := [...]int{1, 2, 3}
	fmt.Println(intArr3)

	intArr4 := [4]int{1, 2, 3}
	fmt.Println(intArr4)

	// accessing array elements
	element := intArr1[0]
	fmt.Println(element)

	// iteration
	for i := 0; i < len(intArr1); i++ {
		item := intArr1[i]
		fmt.Println(item)
	}

	// array bounds
	// run time error -> iteration
}
