// allows functionality to be isolated
// take data as input and return data as output

package main

import "fmt"

func sum(lhs, rhs int) int {
	return lhs + rhs
}

func describe(v1, v2, v3 int) (int, int, int) {
	minVal := min(v1, v2, v3)
	maxVal := max(v1, v2, v3)
	rangeVal := maxVal - minVal

	return minVal, maxVal, rangeVal
}

func main() {
	total := sum(1, 2)
	fmt.Printf("Total: %d\n", total)

	minVal, maxVal, _ := describe(1, 2, 3)
	fmt.Printf("Min: %d Max: %d\n", minVal, maxVal)
}
