package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func even(str int) int {
	Arg := os.Args[1:]
	a := 0
	b := 0
	for i := 1; i <= len(Arg); i++ {
		a = i
		b = a % 2
	}
	return b
}

func isEven(nbr int) bool {
	if even(nbr) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	lengthofArg := 0
	if isEven(lengthofArg) == true {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
