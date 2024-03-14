// each package can have it's own init() function
// all packages will execute init() before main()

package main

import (
	"fmt"
	"regexp"
)

var EmailExpr *regexp.Regexp

func init() {
	compiled, err := regexp.Compile(`.+@.+\..+`)
	if err != nil {
		panic("failed to compile email regexp")
	}

	EmailExpr = compiled
}

func isValidEmail(email string) bool {
	return EmailExpr.Match([]byte(email))
}

func main() {
	fmt.Println(isValidEmail("invalid@example"))
	fmt.Println(isValidEmail("admin@cansky.studio"))
}
