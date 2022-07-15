// match tool checks a string against a pattern.
// If it matches - prints the string, otherwise prints nothing.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) >= 2 {
		fmt.Println(len(strings.Fields(os.Args[1])))
	}
}
