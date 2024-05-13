package hello

import (
	"testing"
)

// Test_withGreetName groups several subtests to verify the Hello function's behavior
// under different input scenarios.
// Sometimes it is useful to group tests around a "thing" and then have subtests describing different scenarios. sA benefit of this approach is you can set up shared code that can be used in the other tests.
func TestHello(t *testing.T) {
	// Subtest: Standard greeting
	// Tests the function with a non-empty name and default language.
	t.Run("Saying Hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	// Subtest: Default name
	// Tests the function with an empty name to check the default greeting.
	t.Run("say 'Hello, world' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	// Subtest: Greeting in Spanish
	// Tests the function with a non-empty name and the Spanish language.
	t.Run("greet in spanish", func(t *testing.T) {
		got := Hello("Darshan", "Spanish")
		want := "Hola, Darshan"
		assertCorrectMessage(t, got, want)
	})

	// Subtest: Greeting in French
	// Tests the function with a non-empty name and the French language.
	t.Run("greet in french", func(t *testing.T) {
		got := Hello("Aadhithya", "French")
		want := "Bonjour, Aadhithya"
		assertCorrectMessage(t, got, want)
	})
}

// assertCorrectMessage is a helper function used to assert that the output from Hello
// matches the expected string. It reports an error if the two strings do not match.
// This function is marked as a helper to indicate that the actual line of failure
// in test output should be reported in the calling function.
func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper() // Marks this function as a helper function.
	if got != want {
		t.Errorf("got %q want %q", got, want) // Reports a failure if the received message does not match the expected message.
	}
}
