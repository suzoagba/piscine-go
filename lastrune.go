package piscine

func LastRune(s string) rune {
	sentence := []rune(s)

	return sentence[len(s)-1]
}
