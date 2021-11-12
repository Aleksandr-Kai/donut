package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	lst := []rune{'.', ',', '-', '~', ':', ';', '=', '!', '*', '#', '$', '@'}
	A := 0.0
	B := 0.0
	z := make([]float64, 7040)
	b := make([]rune, 1760)
	sb := strings.Builder{}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Print("\x1b[2J")
	for {
		for k := 0; k < 1760; k++ {
			b[k] = 32
		}
		for k := 0; k < 7040; k++ {
			z[k] = 0
		}
		for j := 0.0; j < 6.28; j += 0.07 {
			for i := 0.0; i < 6.28; i += 0.02 {
				c := math.Sin(i)
				d := math.Cos(j)
				e := math.Sin(A)
				f := math.Sin(j)
				g := math.Cos(A)
				h := d + 2
				D := 1.0 / (c*h*e + f*g + 5)
				l := math.Cos(i)
				m := math.Cos(B)
				n := math.Sin(B)
				t := c*h*g - f*e
				x := int(40 + 30*D*(l*h*m-t*n))
				y := int(12 + 15*D*(l*h*n+t*m))
				o := x + 80*y
				N := int(8 * ((f*e-c*d*g)*m - c*d*e - f*g - l*d*n))
				if 22 > y && y > 0 && x > 0 && 80 > x && D > z[o] {
					z[o] = D
					if N > 0 {
						b[o] = lst[N]
					} else {
						b[o] = 32
					}
				}
			}
		}
		fmt.Print("\x1b[H")
		sb.Reset()
		for k := 0; k < 1761; k++ {
			if k%80 > 0 {
				sb.WriteRune(b[k])
			} else {
				fmt.Println(sb.String())
				sb.Reset()
			}
			A += 0.00004
			B += 0.00002
		}
		time.Sleep(time.Millisecond * 100)
	}
}
