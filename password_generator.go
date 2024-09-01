package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const (
	defaultLength         = 12
	defaultIncludeSymbols = false
	defaultIncludeNumbers = true
	defaultExcludeSimilar = false
	defaultMinLength      = 8
	defaultMaxLength      = 128
	letters               = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers               = "0123456789"
	symbols               = "!@#$%^&*()-_=+[]{}|;:,.<>/?"
	similarChars          = "lI1oO0"
)

// Function to generate a random password
func generatePassword(length int, includeSymbols, includeNumbers, excludeSimilar bool) (string, error) {
	if length <= 0 {
		return "", errors.New("password length must be greater than zero")
	}

	if length < defaultMinLength || length > defaultMaxLength {
		return "", fmt.Errorf("password length must be between %d and %d", defaultMinLength, defaultMaxLength)
	}

	chars := letters
	if includeNumbers {
		chars += numbers
	}
	if includeSymbols {
		chars += symbols
	}
	if excludeSimilar {
		chars = filterChars(chars, similarChars)
	}

	if len(chars) == 0 {
		return "", errors.New("character set for password is empty")
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

// Filter out similar characters
func filterChars(chars, exclude string) string {
	var result strings.Builder
	for _, c := range chars {
		if !strings.ContainsRune(exclude, c) {
			result.WriteRune(c)
		}
	}
	return result.String()
}

// Main function using cobra for command-line parsing
func main() {
	var length int
	var includeSymbols, includeNumbers, excludeSimilar bool
	var minLength, maxLength int
	var configFile string
	var verbose bool

	var rootCmd = &cobra.Command{
		Use:   "password_generator",
		Short: "Generate a random password with advanced features",
		Long: `This tool generates a random password with advanced options including:
- Length specification
- Inclusion of symbols, numbers, and exclusion of similar characters
- Minimum and maximum length constraints
- Verbose output and configuration file support`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if verbose {
				fmt.Println("Generating password with the following options:")
				fmt.Printf("Length: %d\n", length)
				fmt.Printf("Include Symbols: %v\n", includeSymbols)
				fmt.Printf("Include Numbers: %v\n", includeNumbers)
				fmt.Printf("Exclude Similar Characters: %v\n", excludeSimilar)
				fmt.Printf("Min Length: %d, Max Length: %d\n", minLength, maxLength)
			}

			// Validate length constraints
			if length < minLength || length > maxLength {
				return fmt.Errorf("length must be between %d and %d", minLength, maxLength)
			}

			password, err := generatePassword(length, includeSymbols, includeNumbers, excludeSimilar)
			if err != nil {
				return fmt.Errorf("error generating password: %w", err)
			}

			fmt.Println("Generated password:", password)
			return nil
		},
	}

	rootCmd.Flags().IntVarP(&length, "length", "l", defaultLength, "Length of the password")
	rootCmd.Flags().BoolVarP(&includeSymbols, "symbols", "s", defaultIncludeSymbols, "Include symbols in the password")
	rootCmd.Flags().BoolVarP(&includeNumbers, "numbers", "n", defaultIncludeNumbers, "Include numbers in the password")
	rootCmd.Flags().BoolVarP(&excludeSimilar, "exclude-similar", "x", defaultExcludeSimilar, "Exclude similar characters (like 1, l, I, 0, O)")
	rootCmd.Flags().IntVar(&minLength, "min-length", defaultMinLength, "Minimum length of the password")
	rootCmd.Flags().IntVar(&maxLength, "max-length", defaultMaxLength, "Maximum length of the password")
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "", "Path to a configuration file")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
