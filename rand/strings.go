package rand

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seedRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// StringWithCharset generates a random string from the passed in charset.
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seedRand.Intn(len(charset))]
	}

	return string(b)
}

// String generates a random string from the default charset.
func String(length int) string {
	return StringWithCharset(length, charset)
}
