// two types of concurrent code
// 1. threaded: code run parallels based on number of cpu cores
// 2. asynchronous: code can pause and resume execution, while paused other code can resume

// concurrent code run non deterministically, cannot rely on result being the same each program run
// goroutines allow function to run concurrently

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func counter(x int) {
	for i := 1; i <= x; i++ {
		time.Sleep(100 * time.Millisecond)
		// time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

func sumFile(rd bufio.Reader) int {
	sum := 0
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			return sum
		}

		if err != nil {
			fmt.Println("Error:", err)
		}

		num, err := strconv.Atoi(line[:len(line)-1])
		if err != nil {
			fmt.Println("Error:", err)
		}

		sum += num
	}
}

func main() {
	// go counter(5)
	// fmt.Println("wait for goroutine")
	// time.Sleep(1 * time.Second)
	// fmt.Println("end")

	// counter := 0

	// wait := func(ms time.Duration) {
	// 	time.Sleep(ms * time.Millisecond)
	// 	counter += 1
	// }

	// fmt.Println("Launching goroutines")

	// go wait(100)
	// go wait(900)
	// go wait(1000)
	// go wait(1101)

	// fmt.Println("Launched. Counter = ", counter)
	// time.Sleep(1100 * time.Millisecond)
	// fmt.Println("Waited 1100ms. Counter = ", counter)

	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}
	sum := 0

	for i := 0; i < len(files); i++ {
		file, err := os.Open(files[i])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		rd := bufio.NewReader(file)
		calculate := func() {
			fileSum := sumFile(*rd)
			sum += fileSum
		}

		go calculate()
	}

	time.Sleep(1 * time.Second)
	fmt.Println(sum)
}
