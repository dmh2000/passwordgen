package pwd

import (
	"crypto/rand"
	"strings"
)

func GeneratePassword(length int, includeSymbols bool) string {
	const (
		letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numbers = "0123456789"
		symbols = "!@#$%^&*()-_=+[]{}|;:,.<>?"
	)

	// Create character set
	charset := letters + numbers
	if includeSymbols {
		charset += symbols
	}

	// Generate password
	password := make([]byte, length)
	charsetLen := len(charset)

	for i := 0; i < length; i++ {
		randomBytes := make([]byte, 1)
		_, err := rand.Read(randomBytes)
		if err != nil {
			panic(err)
		}
		password[i] = charset[int(randomBytes[0])%charsetLen]
	}

	return string(password)
}

func FormatPassword(password string) string {
	var builder strings.Builder
	for i := 0; i < len(password); i += 3 {
		if i > 0 {
			builder.WriteString(" ")
		}
		end := i + 3
		if end > len(password) {
			end = len(password)
		}
		builder.WriteString(password[i:end])
	}
	return builder.String()
}
