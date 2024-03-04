
package piscine

import "fmt"


func QuadA(x, y int) {
	for i := 0; i != x; i++ {
		for j := 0; j < x; j++ {
			if i == 0 || i == y-1 || j == 0 || j == x{
				z01.PrintRune("o")
			} else {
				z01.PrintRune("|")
			}
		}
		fmt.Println()
	}
}

func main() {
	QuadA(5, 3)
}