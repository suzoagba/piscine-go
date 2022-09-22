package piscine

func AppendRange(min, max int) []int {
	var myarray []int
	if min > max {
		return myarray
	} else {
		for i := min; i < max; i++ {
			myarray = append(myarray, i)
		}
	}
	return myarray
}
