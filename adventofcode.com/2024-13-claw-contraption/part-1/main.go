package main

import (
	"fmt"
	"io"
)

func main() {
	total := 0
	for {
		pa, pb := 3, 1
		var xa, ya, xb, yb, xc, yc int
		_, err := fmt.Scanf("Button A: X+%d, Y+%d\n", &xa, &ya)
		if err != nil {
			// It feels fishy to expect io.ErrUnexpectedEOF.
			if err == io.ErrUnexpectedEOF {
				break
			}
			panic(err)
		}
		_, err = fmt.Scanf("Button B: X+%d, Y+%d\n", &xb, &yb)
		if err != nil {
			panic(err)
		}
		_, err = fmt.Scanf("Prize: X=%d, Y=%d\n", &xc, &yc)
		if err != nil {
			panic(err)
		}

		minPrice := computeMinPrice(pa, pb, xa, ya, xb, yb, xc, yc)
		if minPrice != nil {
			total += *minPrice
		}

		_, err = fmt.Scanf("\n")
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
	}
	fmt.Println(total)
}

func computeMinPrice(pa, pb, xa, ya, xb, yb, xc, yc int) *int {
	maxA := divCeil(xc, xa)
	maxB := divCeil(xc, xb)

	minPrice := (*int)(nil)
	for a := 0; a <= maxA; a++ {
		for b := 0; b <= maxB; b++ {
			if xa*a+xb*b == xc && ya*a+yb*b == yc {
				p := pa*a + pb*b
				if minPrice == nil {
					minPrice = new(int)
					*minPrice = p
				} else {
					*minPrice = min(*minPrice, p)
				}
			}
		}
	}

	return minPrice
}

func divCeil(a, b int) int {
	c := a / b
	if a%b != 0 {
		c++
	}
	return c
}
