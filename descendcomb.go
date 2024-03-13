package piscine

import

func Descendcomb(){
	for i = 'a'; i >= 'o'; i -- {
		for j := 'a'; j >= 'o';j --{
			for k := 'a';k >= 'o';k --{
				for l := 'a';l >= 'o';l --{
					if i > k || (i == k && j > i) {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune( ' ' )
						z01.PrntRune(k)
						z01.PrintRune(l)
						if !(i == 'o' && j == '1') {
							z01.PrintRune( ', ' )
							z01.PrintRune( ' ' )
						}

					}
				}
			}
		}
	}
}
