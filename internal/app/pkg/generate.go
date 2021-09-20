package pkg

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

const defaultCodeLength = 8

var chars = []rune(
	"abcdefghijkmnpqrstuvwxyz" +
		"123456789",
)

// NewStringGenerator is a random string generator.
func NewGeneratedString() (string, error) {
	value, err := generateRandomString(defaultCodeLength)
	if err != nil {
		return "", err
	}

	return value, nil
}

func generateRandomString(length int) (string, error) {
	if length < 1 {
		return "", errors.New("invalid code length provided")
	}

	rand.Seed(time.Now().UnixNano())
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String(), nil
}
