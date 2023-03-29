package main

import (
	"math/rand"
	"strings"
)

func alphabetGenerator() string {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := strings.ToUpper(lower)
	numbers := "0123456789"
	special := "!@#$%^&*()"
	return lower + upper + numbers + special
}

func PasswordGenerator(length int) string {
	alphabet := alphabetGenerator()
	n := len(alphabet)

	var password string

	for i := 0; i < length; i++ {
		index := rand.Intn(n)
		password += string(alphabet[index])
	}

	return password
}
