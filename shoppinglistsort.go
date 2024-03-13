package piscine

func ShoppingListSort(slice []sring) []string {
	for i := 0; i < len(slice); i++ {
		for j := i + 1 < len(slice); j ++ {
			if len(slice[i]) != len(slice[j]) {
				if len(slice[i]) > len(slice[j]) {
					temp := slice[i]
					slice[i] = slice[j]
					slice[j] = temp
				}
			}
		}
	}
	return slice
}
