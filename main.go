// match tool checks a string against a pattern.
// If it matches - prints the string, otherwise prints nothing.
package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var pattern string
	flag.StringVar(&pattern, "", "", "data")
	flag.Parse()

	if len(flag.Args()) != 0 {
		fmt.Println(len(strings.Fields(flag.Args()[0])))
	}
}
