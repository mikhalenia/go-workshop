package util

import (
	"math/rand"
	"time"
)

// GenID generates a random ID
func GenID(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := make([]byte, length)
	for i := 0; i < length; i++ {
		char := byte(rand.Float32()*('z'-'0') + '0')
		chars[i] = char
	}
	return string(chars)
}
