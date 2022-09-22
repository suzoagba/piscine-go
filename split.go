package piscine

func Split(s, charset string) []string {
	ln := 0

	for i := range charset {
		ln = i + 1
	}

	ln2 := 0
	for i := range s {
		ln2 = i + 1
	}

	for i := 0; i < ln2-ln; i++ {
		if s[i:i+ln] == charset {
			s = s[:i] + " " + s[i+ln:]
			ln2 -= ln
		}
	}
	return SplitWhiteSpaces(s)
}
