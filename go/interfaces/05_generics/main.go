// allow function to handle multiple types of data
// generics are defined using inteface (called constraints)

package main

import (
	"fmt"

	"github.com/naufalihsan/academy/constraints"
)

// builtin constraints
func IsEqual[T comparable](a, b T) bool {
	return a == b
}

// create a constraint
type Integers32 interface {
	// (~) approximation to allow type aliases
	~uint32 | int32
}

func SumNumbers[T Integers32](arr []T) T {
	var sum T
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

// generic structure
type MyArray[T constraints.Ordered] struct {
	arr []T
}

func (m *MyArray[T]) Max() T {
	max := m.arr[0]
	for i := 0; i < len(m.arr); i++ {
		if m.arr[i] > max {
			max = m.arr[i]
		}
	}
	return max
}

func main() {
	fmt.Println(IsEqual(1, 1))
	fmt.Println(IsEqual("foo", "bar"))
	fmt.Println(IsEqual('a', 'a'))
	fmt.Println(IsEqual[uint8](0, 1))

	arr1 := []uint32{0, 1, 2}
	fmt.Println(SumNumbers(arr1))

	arr2 := []int32{3, 4, 5}
	fmt.Println(SumNumbers(arr2))

	type MyInt uint32
	arr3 := []MyInt{MyInt(6), MyInt(7), MyInt(8)}
	fmt.Println(SumNumbers(arr3))

	arr4 := MyArray[int]{arr: []int{1, 2, 8, 3, 4, 5}}
	fmt.Println(arr4.Max())
}
