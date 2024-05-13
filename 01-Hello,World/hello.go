// Package hello provides a multilingual greeting function.
// It demonstrates basic principles of Test-Driven Development (TDD) by
// allowing greetings in English, Spanish, and French.
package hello

import (
	"fmt"
)

// Constants for supported languages.
const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// Hello returns a greeting string addressed to the provided name in the specified language.
// If no name is provided, it defaults to "world".
// Supported languages are English, Spanish, and French.
// If the language provided is not supported, it defaults to English.
//
// The function is designed with TDD principles:
// 1. Write a failing test: Start by writing a test that fails because the function it tests does not yet exist.
// 2. Make the test pass: Write the minimum amount of code necessary for the test to pass.
// 3. Refactor: Clean up the new code, ensuring it fits well with the existing design.
//
// Example usage:
//
//	Hello("Alice", "French") // returns "Bonjour, Alice"
//	Hello("", "")           // returns "Hello, world"
func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

// greetingPrefix determines the greeting message prefix based on the provided language.
// This helper function helps encapsulate the logic for prefix selection,
// making the main Hello function cleaner and easier to manage.
// It returns a greeting prefix for supported languages or defaults to English if unspecified or unsupported.
// We will be using switch instead of if condition to support more language condition in the future.
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// main serves as an entry point of the application, demonstrating the use of the Hello function.
func main() {
	fmt.Println(Hello("World!", ""))
}
