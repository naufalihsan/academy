// commonly data structure that stores data in key value pairs
// high performance when the key is known
// unordered data

package main

import "fmt"

func main() {
	// create a map
	initMap := make(map[int]bool)
	fmt.Println(initMap)

	itemMap := map[int]bool{
		1: true,
		2: false,
		3: true,
	}
	fmt.Println(itemMap)

	// map operations
	itemMap[4] = false // insert

	item := itemMap[4] // read
	fmt.Println(item)

	delete(itemMap, 4) // delete

	item, found := itemMap[4] // check existence
	if !found {
		fmt.Println("item not found")
	}

	// iteration (unordered)
	for key, value := range itemMap {
		fmt.Println(key, value)
	}
}
