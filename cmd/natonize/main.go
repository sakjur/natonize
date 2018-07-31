package main

import (
	"github.com/sakjur/natonize/pkg/natonize"
	"strings"
	"fmt"
	"flag"
	)

func main() {
	var output string

	reverseNato := flag.Bool("reverse", false, "Reverse a string with NATO values")
	flag.Parse()

	args := flag.Args()
	input := strings.Join(args, " ")

	if input == "" {
		fmt.Println("Usage: natonize [--reverse] <term to natonize>")
	}

	if *reverseNato {
		output = natonize.FromNatoAlphabet(input)
	} else {
		output = natonize.ToNatoAlphabet(input)
	}

	fmt.Println(output)
}
