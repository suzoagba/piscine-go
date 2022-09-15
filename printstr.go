package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	Astr := []rune(s)
	for _, v := range Astr {
		z01.PrintRune(v)
	}
}
