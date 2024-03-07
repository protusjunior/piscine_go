package piscine

func Index(s string, toFind string) int {
	for i := 0; i <= len(s)-len(toFind); i++ {
		k := s[i : i+len(toFind)]
		if k == toFind {
			return i
		}
	}
	return -1
}
