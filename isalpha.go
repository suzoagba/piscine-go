package piscine

func IsAlpha(s string) bool {
	c := []rune(s)
	cd := 0

	for i := range s {
		cd = i
	}
	for i := 0; i <= cd; i++ {
		if (c[i] < 'a' || c[i] > 'z') &&
			(c[i] < 'A' || c[i] > 'Z') &&
			(c[i] < '0' || c[i] > '9') {
			return false
		}
	}
	return true
}
