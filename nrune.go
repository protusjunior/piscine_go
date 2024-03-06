package piscine

func NRune(s string, n int) rune {
	if n < 1 || n > len(s) {
		return '\x00'
	}
	
	return rune([]rune(s)[n])
}