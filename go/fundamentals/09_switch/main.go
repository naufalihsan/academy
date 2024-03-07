// switch is used to easily check multiple condition

package main

import "fmt"

func getCurrentBrowser() string {
	return "edge"
}

func main() {
	browser := "chrome"

	switch browser {
	case "chrome":
		fmt.Println(1)
	case "mozilla":
		fmt.Println(2)
	default:
		fmt.Println("browser not available:", browser)
	}

	// conditional cases
	switch currentBrowser := getCurrentBrowser(); {
	case currentBrowser == "chrome":
		fmt.Println(1)
	case currentBrowser == "mozilla":
		fmt.Println(2)
	default:
		fmt.Println("browser not available:", currentBrowser)
	}

	// case list
	char := 'a'
	switch char {
	case ' ':
		// nothing happened
	case 'a', 'i', 'u', 'e', 'o':
		fmt.Println("vowel")
		fallthrough // will execute the next case
	case 'A', 'I', 'U', 'E', 'O':
		fmt.Println("vowel too")
	default:
		fmt.Println("not vowel")
	}
}
