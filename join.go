package piscine

func Join(str []string, sep string) string {
	var nwStr string
	for i, c := range str {
		nwStr += c
		if i < len(str)-1 {
			nwStr += sep
		}
	}
	return nwStr
}
