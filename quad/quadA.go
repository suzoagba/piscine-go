package main

import (
	"github.com/01-edu/z01"
)

func Raid1a(x, y int) {

	if x > 0 && y > 0 {

		if x == 1 && y == 1 {
			z01.PrintRune('A')
			z01.PrintRune('\n')
			return
		} else if x == 1 {
			z01.PrintRune('A')
			for t := 1; t < y-1; t++ {
				if t <= 0 {
					break
				} else {
					z01.PrintRune('\n')
					z01.PrintRune('B')

				}
				if y == 1 {
					z01.PrintRune('\n')
					z01.PrintRune('A')
					return
				}
			}
			z01.PrintRune('\n')
			z01.PrintRune('c')
			z01.PrintRune('\n')

		} else if y == 1 {
			z01.PrintRune('A')
			for i := 1; i < x-1; i++ {
				z01.PrintRune('B')
			}
			z01.PrintRune('A')
			z01.PrintRune('\n')
			return
		} else {
			z01.PrintRune('A')
			for i := 1; i < x-1; i++ {
				z01.PrintRune('B')
			}
			z01.PrintRune('A')
			z01.PrintRune('\n')

			for t := 1; t < y-1; t++ {
				z01.PrintRune('B')

				for j := 1; j < x-1; j++ {
					z01.PrintRune(' ')
				}
				z01.PrintRune('B')
				z01.PrintRune('\n')

			}

			z01.PrintRune('C')

			for i := 1; i < x-1; i++ {
				z01.PrintRune('B') 
			}
			z01.PrintRune('C')
			z01.PrintRune('\n')

		}
	}
}

func main() {
	Raid1a(5, 3)
}
