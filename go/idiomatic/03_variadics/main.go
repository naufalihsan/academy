// make a function that accept many number of parameters

package main

import "fmt"

func sum(nums ...int) int {
	s := 0

	for _, n := range nums {
		s += n
	}

	return s
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := append(a, b...)

	fmt.Println(sum(c...))

	fmt.Println(sum(1, 2, 3, 4, 5, 6))
}
