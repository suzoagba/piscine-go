package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	m := 96
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "--upper" && i == 1 {
			m -= 32
			continue
		}
		if basicAtoi2(os.Args[i]) == 0 {
			z01.PrintRune(' ')
		} else {
			if basicAtoi2(os.Args[i]) <= 26 {
				z01.PrintRune(rune(m + basicAtoi2(os.Args[i])))
			} else {
				z01.PrintRune(' ')
			}
		}
	}
	if len(os.Args) > 1 {
		z01.PrintRune('\n')
	}
}

func basicAtoi2(s string) int {
	sum := 0
	for _, str := range s {
		sum *= 10
		sum += int(rune(str)) - 48
		if rune(str) > rune('9') || rune(str) < rune('0') {
			sum = 0
			break
		}
	}
	return sum
}
