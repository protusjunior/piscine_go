package piscine

func AlphaCount(s string) int {
	var count int
	for _, v := range s {
		if ('a' <= v && v <='z') || ('A' <= v & v <='z') {
			count++
		}		
		return count
	}
}
