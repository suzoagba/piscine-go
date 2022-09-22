package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	} else {
		myarray := make([]int, max-min)
		for i := 0; i < max-min; i++ {
			myarray[i] = i + min
		}
		return myarray
	}
}
