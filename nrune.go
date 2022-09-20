package piscine

func NRune(s string, n int) rune {
	sentence := []rune(s)
	for index := range sentence {
		if index == n-1 {
			return sentence[n-1]
		}
	}
	return 0
}
