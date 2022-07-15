// match tool checks a string against a pattern.
// If it matches - prints the string, otherwise prints nothing.
package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	//rdr := bufio.NewReader(os.Stdin)
	//tr, _ := rdr.ReadString('\n')

	var pattern string
	flag.StringVar(&pattern, "", "", "data")
	flag.Parse()

	fmt.Println(len(strings.Fields(flag.Args()[0])))
}
