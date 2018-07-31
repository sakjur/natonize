package main

import (
	"os"
	"github.com/sakjur/natonize/pkg/natonize"
	"strings"
	"fmt"
)

func main() {
	args := os.Args[1:]

	fmt.Println(natonize.ToNatoAlphabet(strings.Join(args, " ")))
}
