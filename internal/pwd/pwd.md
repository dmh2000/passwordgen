# Password Generator Package

A Go package that provides secure password generation and formatting utilities.

## Features

- Generates cryptographically secure random passwords
- Minimum password length of 24 characters for enhanced security
- Optional inclusion of special symbols
- Password formatting utility for better readability

## Functions

### `GeneratePassword(length int, includeSymbols bool) (string, error)`

Generates a random password with the following characteristics:
- Always includes letters (a-z, A-Z) and numbers (0-9)
- Optionally includes special symbols (!@#$%^&*()-_=+[]{}|;:,.<>?)
- Returns an error if length is less than 24 characters
- Uses crypto/rand for secure random generation

### `FormatPassword(password string) (string, error)`

Formats a password string by:
- Adding spaces every 3 characters
- Making the password more readable and easier to communicate

## Usage Example

```go
password, err := pwd.GeneratePassword(24, true)
if err != nil {
    log.Fatal(err)
}

formatted, err := pwd.FormatPassword(password)
if err != nil {
    log.Fatal(err)
}
```

## Security Note

This package enforces a minimum password length of 24 characters to provide adequate security against brute-force attacks. It uses Go's crypto/rand package for cryptographically secure random number generation.
