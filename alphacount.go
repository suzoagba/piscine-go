package piscine

func AlphaCount(s string) int {
	count := 0
	sentence := []rune(s)

	for i := 0; i < len(s); i++ {
		if sentence[i] >= 'a' && sentence[i] <= 'z' || sentence[i] >= 'A' && sentence[i] <= 'Z' {
			count++
		}
	}
	return count
}
