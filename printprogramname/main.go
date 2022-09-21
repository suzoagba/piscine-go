package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	runeargs := []rune(os.Args[0])            // take the path only from os.Args and make it a rune slice
	runeprint := []rune{}                     // make the rune list to print
	for i := len(runeargs) - 1; i >= 0; i-- { // loop arround each ellement in the rune slice backwards
		if '.' == runeargs[i] { // is the element reach . then empty the rune to print
			runeprint = []rune{} // reinitialize rune to print
		} else if '/' == runeargs[i] { // if it reach / stop
			break
		} else { // append the char in the rune to print
			runeprint = append(runeprint, runeargs[i])
		}
	}
	for i := len(runeprint) - 1; i >= 0; i-- { // print the rune to print backwards since last to come first to go
		z01.PrintRune(runeprint[i])
	}
	z01.PrintRune('\n') // print /n needed in the end
}
