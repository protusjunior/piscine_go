package piscine

func Map(f func(int) bool, a []int) []bool {
	boolVal := make([]bool, len(a))
	for i, s := range a {
		boolVal[i] = f(s)
	}
	return boolVal
}
