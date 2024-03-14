package main

import "regexp"

func IsValidEmail(email string) bool {
	re, err := regexp.Compile(`.+@.+\..+`)
	if err != nil {
		panic("failed to compile email regexp")
	}
	return re.Match([]byte(email))
}
