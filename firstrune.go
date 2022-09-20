package piscine

func FirstRune(s string) rune {
	sentence := []rune(s)

	for index := range sentence {
		if index == 0 {
			return sentence[0]
		}
	}
	return 0
}
