package main

import "github.com/01-edu/z01"

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

var str string = "x = 42, y = 21"

func main() {
	for _, st := range str {
		z01.PrintRune(st)
	}
	z01.PrintRune('\n')
}
