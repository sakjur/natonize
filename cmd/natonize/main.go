package main

import (
	"flag"
	"fmt"
	"github.com/sakjur/natonize/pkg/natonize"
	"strings"
)

func main() {
	var output string

	reverseNato := flag.Bool("invert", false, "Invert a string with NATO values")
	flag.Parse()

	args := flag.Args()
	input := strings.Join(args, " ")

	if input == "" {
		fmt.Println("Usage: natonize [--invert] <term to natonize>")
		return
	}

	if *reverseNato {
		output = natonize.FromNatoAlphabet(input)
	} else {
		output = natonize.ToNatoAlphabet(input)
	}

	fmt.Println(output)
}
