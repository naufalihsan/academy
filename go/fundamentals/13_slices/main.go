// enable "view" into an array
// views are dynamic and not fixed size
// function can accept slices as a function parameters

package main

import "fmt"

func iter(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

func main() {
	// creating a slice
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	// accessing element in slice
	element := slice[0]
	fmt.Println(element)

	// slice syntax
	// start index (inclusive) : end index (exclusive)
	arr := [...]int{1, 2, 3, 4}

	slice1 := arr[:]
	fmt.Println(slice1)

	slice2 := arr[2:]
	fmt.Println(slice2)

	slice3 := arr[1:4]
	fmt.Println(slice3)

	// dynamic array
	// extends slice with another slice
	part1 := []int{1, 2, 3}
	part2 := []int{4, 5, 6}
	combined := append(part1, part2...)
	fmt.Println(combined)
	// add additional elements
	combined = append(part1, 4, 5, 6)
	fmt.Println(combined)

	// pre-allocation a slice
	alloc := make([]int, 10)
	fmt.Println(alloc)

	// slices to functions
	small := []int{1, 2}
	iter(small)

	large := []int{3, 4, 5, 6, 7, 8}
	iter(large)

	// multidimensional slices
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix[1][1])
}
