package builtinstring

import (
	"fmt"
	"strings"
)

// This calls built in compare function from strings package of golang
func main() {
	fmt.Println(strings.Compare("a", "a"))
}
