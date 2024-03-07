package main

import (
	"github.com/fatih/color"
	"github.com/naufalihsan/academy/internal/random"
)

func main() {
	cli := color.New(color.FgGreen)
	cli.Printf("Random number: %d\n", random.Number())
}
