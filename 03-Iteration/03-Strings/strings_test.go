package builtinstring_test

import (
	"fmt"
	"strings"
	"testing"
)

func TestCompare(t *testing.T) {
	t.Run("compare two same strings", func(t *testing.T) {
		got := strings.Compare("hi", "hi")
		want := 0
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("compare two strings, with first parameter alphabetically bigger than another", func(t *testing.T) {
		got := strings.Compare("b", "a")
		want := 1
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("compare two strings, with first parameter alphabetically smaller than another", func(t *testing.T) {
		got := strings.Compare("a", "b")
		want := -1
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func BenchmarkCompare(t *testing.B) {
	for i := 0; i < (t.N); i++ {
		strings.Compare("a", "a")
	}
}

func ExampleCompare() {
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("b", "a"))
	// Output: 0
	// -1
	// 1
}
