package hello

import (
	"testing"
)

func Test_withGreetName(t *testing.T) {
	t.Run("Saying Hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("greet in spanish", func(t *testing.T) {
		got := Hello("Darshan", "Spanish")
		want := "Hola, Darshan"
		assertCorrectMessage(t, got, want)
	})
	t.Run("greet in french", func(t *testing.T) {
		got := Hello("Aadhithya", "French")
		want := "Bonjour, Aadhithya"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
