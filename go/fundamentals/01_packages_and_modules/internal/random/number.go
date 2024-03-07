package random

import (
	"math/rand"
)

func Number() int {
	return rand.Intn(100)
}
