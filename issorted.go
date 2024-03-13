package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	asc := true
	desc := true

	for i := 1; i < len(a); i++ {
		if !(f(a[i-1], a[i]) >= 0) {
			desc = false
			break
		}
	}
	for i := 1; i < len(a); i++ {
		if !(f(a[i-1], a[i]) <= 0) {
			asc = false
			break
		}
	}
	return asc || desc
}
