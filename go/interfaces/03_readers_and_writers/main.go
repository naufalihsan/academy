// readers and writers are interfaces that allow reading from & writing to I/O sources

package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func basicRead(text string) {
	reader := strings.NewReader(text)

	var result strings.Builder
	buffer := make([]byte, 4)

	for {
		n, err := reader.Read(buffer)
		chunk := buffer[:n]
		result.Write(chunk)
		fmt.Printf("Read %v bytes: %c\n", n, chunk)
		if err == io.EOF {
			break
		}
	}

	fmt.Println(result.String())
}

func bufioRead(text string) {
	reader := strings.NewReader(text)
	buffered := bufio.NewReader(reader)
	newString, err := buffered.ReadString('\n')
	if err == io.EOF {
		fmt.Println(newString)
	} else {
		fmt.Println("something went wrong")
	}
}

func bufioReadCsv(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic("file not found")
	}

	reader := csv.NewReader(bufio.NewReader(f))

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Printf("Question: %v Answer %v\n", row[0], row[1])
	}
}

func bufioScan() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0, 5)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
		return
	}

	for _, line := range lines {
		fmt.Printf("Line: %v\n", line)
	}
}

func basicWrite(text string) {
	buffer := bytes.NewBufferString("")
	n, err := buffer.WriteString(text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Wrote %v bytes: %c\n", n, buffer)
}

func main() {
	filename := "example.csv"

	basicRead(filename)
	basicWrite(filename)
	bufioReadCsv(filename)
}
