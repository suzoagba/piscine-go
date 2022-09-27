package piscine

func SortWordArr(a []string) {
	c := 0

	for i := range a {
		c = i
	}

	for i := 0; i <= c; i++ {
		for j := 0; j <= c; j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
