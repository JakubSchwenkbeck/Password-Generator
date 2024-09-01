package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	defaultLength         = 12
	defaultIncludeSymbols = false
	letters               = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers               = "0123456789"
	symbols               = "!@#$%^&*()-_=+[]{}|;:,.<>/?"
)

// Function to generate a random password
func generatePassword(length int, includeSymbols bool) (string, error) {
	if length <= 0 {
		return "", errors.New("password length must be greater than zero")
	}

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

// Main function using cobra for command-line parsing
func main() {
	var length int
	var includeSymbols bool

	var rootCmd = &cobra.Command{
		Use:   "password_generator",
		Short: "Generate a random password",
		Long:  `This tool generates a random password with optional symbols.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			password, err := generatePassword(length, includeSymbols)
			if err != nil {
				return fmt.Errorf("error generating password: %w", err)
			}

			fmt.Println("Generated password:", password)
			return nil
		},
	}

	rootCmd.Flags().IntVarP(&length, "length", "l", defaultLength, "Length of the password")
	rootCmd.Flags().BoolVarP(&includeSymbols, "symbols", "s", defaultIncludeSymbols, "Include symbols in the password")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
