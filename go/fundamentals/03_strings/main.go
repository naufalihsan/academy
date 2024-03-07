// textual data in go use UTF-8 encoding using code pages
// each symbol in code pages is called a code point

// text is represented by rune type (similar to char)
// rune is alias for int32 (number represent symbol)

// string is the data type for storing multiple runes

package main

import "fmt"

func main() {
	var runeA rune = 'A'
	fmt.Println(runeA)

	var runeHorse rune = '𓃗'
	fmt.Println(runeHorse)

	sentence := "A horse 𓃗"
	fmt.Println(sentence)

	rawLiteral := `Let's code in \"Golang!"\n`
	fmt.Println(rawLiteral)
}
