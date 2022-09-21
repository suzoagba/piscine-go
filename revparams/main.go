package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	c := 0

	for i := range arg {
		c = i
	}

	for i := c; i >= 1; i-- {
		for _, w := range arg[i] {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}
