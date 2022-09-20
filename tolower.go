package piscine

func ToLower(s string) string {
	my_arr := []rune(s)
	for index, word := range my_arr {
		if word >= 'A' && word <= 'Z' {
			my_arr[index] = word + 32
		}
	}
	return string(my_arr)
}
