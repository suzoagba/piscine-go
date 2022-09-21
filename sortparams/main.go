package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	arg2 := os.Args
	c := 0

	for i := range arg {
		c = i
	}
	for i := 1; i <= c; i++ {
		for j := 1; j <= c; j++ {
			if arg[i] < arg2[j] {
				arg2[j], arg[i] = arg[i], arg2[j]
			}
		}
	}

	for i := 1; i <= c; i++ {
		for _, w := range arg2[i] {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}
