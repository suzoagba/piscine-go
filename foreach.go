package piscine

import "github.com/01-edu/z01"

func ForEach(f func(int), a []int) {
	for i := range a {
		f(a[i])
	}
	return
	
 	z01.PrintRune('\n')
	

}
