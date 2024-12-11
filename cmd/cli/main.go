package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	pwd "sqirvy.xyz/passwords/internal/pwd"
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
		password := pwd.GeneratePassword(passwordLength, *symbols || *symbolsLong)
		fmt.Println()
		fmt.Println(strings.Repeat("-", 80))
		fmt.Println(password)
		fmt.Println(pwd.FormatPassword(password))
	}
}
