package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else if 1 < power {
		nb = nb * RecursivePower(nb, power-1)
	}
	return nb
}
