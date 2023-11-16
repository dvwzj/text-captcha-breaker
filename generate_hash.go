package tcb

import (
	"math/rand"
	"strconv"
)

// GenerateHash generates a random hash
func GenerateHash() string {
	// hash = Math.random().toString(36).substring(2) in javascript (11 characters)
	// hash = strconv.FormatInt(rand.Int63(), 36)[2:] in go (10 characters)
	// hash = strconv.FormatInt(rand.Int63(), 36) in go (12 characters)
	// we need 11 characters
	return strconv.FormatInt(rand.Int63(), 36)[:11]
}
