### Convert Array to Slice (Vice Versa)
Arrays have a fixed size, while slices are dynamic and can change in length

#### Array to Slice
```bash
package main

import (
    "fmt"
)

func main() {
    // Array
    arr := [5]int{1, 2, 3, 4, 5}

    // Convert array to slice
    slice := arr[:]  // slice of the entire array
    // or
    // slice := arr[1:4]  // slice of a portion of the array

    // Print the slice
    fmt.Println("Slice:", slice)
}
```

#### Slice to Array
```bash
package main

import (
    "fmt"
)

func main() {
    // Slice
    slice := []int{1, 2, 3, 4, 5}

    // Create an array with the same length as the slice
    arr := make([]int, len(slice))

    // Copy elements from slice to array
    copy(arr, slice)

    // Print the array
    fmt.Println("Array:", arr)
}
```