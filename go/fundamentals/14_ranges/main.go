package main

import "fmt"

func main() {
	words := []string{"I", "love", "coding"}

	for i, element := range words {
		fmt.Println(i, element)
		for _, char := range element {
			fmt.Printf("%q\n", char)
		}
	}
}
