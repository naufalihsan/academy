// channels offer bidirectional communications
// conceptually the same as a two ended pipe
// write data in one and read data out the other (sending and receiving)

// channels enables goroutines to communicate
// channels can be buffered (specified capacity) or unbuffered (block when sending until a reader is available)
// messages on a channel are FIFO ordering

package main

import "fmt"

func main() {
	// channel := make(chan int)

	// // send to channel
	// go func() { channel <- 1 }()
	// go func() { channel <- 2 }()
	// go func() { channel <- 3 }()

	// // receive from channel
	// first := <-channel
	// second := <-channel
	// third := <-channel

	// fmt.Println(first, second, third)

	// buffered channel
	channel := make(chan int, 2)
	channel <- 1
	channel <- 2
	// channel <- 3
	go func() { channel <- 4 }()

	first := <-channel
	second := <-channel
	third := <-channel

	fmt.Println(first, second, third)
}
