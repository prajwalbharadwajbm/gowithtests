package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	t.Run("Repeat a five times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 5))
	// Output: aaaaa
}

// NOTE by default Benchmarks are run sequentially.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}

}
