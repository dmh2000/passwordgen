package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	minLength     = 24
	defaultLength = 24
	defaultCount  = 1
)

func main() {
	// Define command line flags
	length := flag.Int("l", defaultLength, "Length of password (minimum 20)")
	lengthLong := flag.Int("length", defaultLength, "Length of password (minimum 20)")
	count := flag.Int("c", defaultCount, "Number of passwords to generate")
	countLong := flag.Int("count", defaultCount, "Number of passwords to generate")
	symbols := flag.Bool("s", false, "Include symbols in password")
	symbolsLong := flag.Bool("symbols", false, "Include symbols in password")
	help := flag.Bool("h", false, "Show help message")
	helpLong := flag.Bool("help", false, "Show help message")

	flag.Parse()

	// Handle help flag
	if *help || *helpLong {
		flag.Usage()
		os.Exit(0)
	}

	// Use the long version if it's set, otherwise use short version
	passwordLength := *length
	if flag.Lookup("length").Value.String() != fmt.Sprint(defaultLength) {
		passwordLength = *lengthLong
	}

	passwordCount := *count
	if flag.Lookup("count").Value.String() != fmt.Sprint(defaultCount) {
		passwordCount = *countLong
	}

	// Validate minimum length
	if passwordLength < minLength {
		fmt.Fprintf(os.Stderr, "Password length must be at least %d characters\n", minLength)
		os.Exit(1)
	}

	// Generate passwords
	for i := 0; i < passwordCount; i++ {
		password := generatePassword(passwordLength, *symbols || *symbolsLong)
		fmt.Println(password)
		fmt.Println(formatPassword(password))
		fmt.Println()
	}
}

func generatePassword(length int, includeSymbols bool) string {
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

func formatPassword(password string) string {
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
