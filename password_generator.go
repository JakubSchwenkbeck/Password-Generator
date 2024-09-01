package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

// Function to generate a random password
func generatePassword(length int, includeSymbols bool) (string, error) {
	const (
		letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numbers = "0123456789"
		symbols = "!@#$%^&*()-_=+[]{}|;:,.<>/?"
	)

	chars := letters + numbers
	if includeSymbols {
		chars += symbols
	}

	password := make([]byte, length)
	_, err := rand.Read(password)
	if err != nil {
		return "", err
	}

	for i := range password {
		password[i] = chars[int(password[i])%len(chars)]
	}

	return string(password), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: password_generator <length> [includeSymbols]")
		return
	}

	length := 12
	fmt.Sscanf(os.Args[1], "%d", &length)

	includeSymbols := false
	if len(os.Args) > 2 {
		includeSymbols = os.Args[2] == "true"
	}

	password, err := generatePassword(length, includeSymbols)
	if err != nil {
		fmt.Println("Error generating password:", err)
		return
	}

	fmt.Println("Generated password:", password)
}
