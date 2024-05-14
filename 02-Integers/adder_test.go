package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// This function can be used as Example in godoc and allow users to run and check the functions usage.
// As it executes the example, the testing framework captures data written to standard output and then compares the output against the example’s “Output:” comment.
// The test passes if the test’s output matches its output comment.
// If the Output comment is removed then the function is just compiled but not executed.
func ExampleAdd() {
	fmt.Println(Add(2, 2))
	// Output: 4
}

// It is a good practice to have this file in package 'integers_test'.
