package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, c := range tab {
		if f(c) == true {
			count++
		}
	}
	return count
}
