package hello

import "fmt"

// Printing Hello World Directly to stdout will be outside the domain, thus we would not be able to test this.
/* Hence learn-go-with-tests suggest us to use a function to send in a string to Println(outside world), as we have control over the string and also achieving the goal of printing it to outside world(stdout) */
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("World!"))
}
