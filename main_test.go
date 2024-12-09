package main

import (
	"strings"
	"testing"
	"unicode"
)

func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		name           string
		length         int
		includeSymbols bool
		wantLength     int
		wantSymbols    bool
	}{
		{
			name:           "Default length without symbols",
			length:         24,
			includeSymbols: false,
			wantLength:     24,
			wantSymbols:    false,
		},
		{
			name:           "Minimum length with symbols",
			length:         20,
			includeSymbols: true,
			wantLength:     20,
			wantSymbols:    true,
		},
		{
			name:           "Long password without symbols",
			length:         50,
			includeSymbols: false,
			wantLength:     50,
			wantSymbols:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generatePassword(tt.length, tt.includeSymbols)

			// Check length
			if len(got) != tt.wantLength {
				t.Errorf("generatePassword() length = %v, want %v", len(got), tt.wantLength)
			}

			// Check for symbols
			hasSymbols := strings.ContainsAny(got, "!@#$%^&*()-_=+[]{}|;:,.<>?")
			if hasSymbols != tt.wantSymbols {
				t.Errorf("generatePassword() hasSymbols = %v, want %v", hasSymbols, tt.wantSymbols)
			}

			// Check for letters and numbers
			hasLetters := false
			hasNumbers := false
			for _, char := range got {
				if unicode.IsLetter(char) {
					hasLetters = true
				}
				if unicode.IsNumber(char) {
					hasNumbers = true
				}
			}
			if !hasLetters {
				t.Error("generatePassword() should contain letters")
			}
			if !hasNumbers {
				t.Error("generatePassword() should contain numbers")
			}
		})
	}
}

func TestFormatPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     string
	}{
		{
			name:     "Empty password",
			password: "",
			want:     "",
		},
		{
			name:     "Password length < 3",
			password: "ab",
			want:     "ab",
		},
		{
			name:     "Password length = 3",
			password: "abc",
			want:     "abc",
		},
		{
			name:     "Password length > 3",
			password: "abcdef",
			want:     "abc def",
		},
		{
			name:     "Password length not divisible by 3",
			password: "abcdefgh",
			want:     "abc def gh",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatPassword(tt.password); got != tt.want {
				t.Errorf("formatPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
