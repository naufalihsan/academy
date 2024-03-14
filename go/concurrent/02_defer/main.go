// defer useful to run operations after function complete
// can be used to execute code ater a function runs (clean up resource)

package main

import (
	"fmt"
	"os"
)

func one() {
	fmt.Println(1)
}

func two() {
	fmt.Println(2)
}

func three() {
	fmt.Println(3)
}

func run() {
	fmt.Println("start")

	// stacked (LIFO)
	defer one()
	defer two()
	defer three()

	fmt.Println("end")
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	buffer := make([]byte, 30)
	bytes, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", buffer[:bytes])
}

func main() {
	run()

	filename := "example.txt"
	readFile(filename)
}
