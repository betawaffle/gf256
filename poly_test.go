package gf256

import (
	"fmt"
	"testing"
)

func TestMinimal(t *testing.T) {
	if !testing.Verbose() {
		return
	}

	poly(1<<8 | 1<<4 | 1<<3 | 1<<1 | 1<<0).findOthers()
}

// func TestPoly(t *testing.T) {
// 	p := order | uint64(r)
//
// 	if eval(p, 0) == 0 || eval(p, 1) == 0 {
// 		t.Fatalf("(x^%d + %s) is not irreducible", order, r)
// 	}
// }

// func eval(p uint64, x Element) Element {
// 	if x == 0 {
// 		return Element(p & 1)
// 	}
//
// 	y := Element(p >> 63)
// 	i := uint(63)
//
// 	for i > 0 {
// 		i--
// 		y = Mul(y, x) ^ Element(p>>i&1)
// 	}
// 	return y
// }

type poly uint64

func (r poly) degree() (n uint8) {
	for r > 1 {
		r >>= 1
		n++
	}
	return
}

func (r poly) square(a poly) poly {
	return r.mul(a, a)
}

func (r poly) mul(a, b poly) (p poly) {
	order := poly(1) << r.degree()
	for i := uint(0); b > 0; i++ {
		if b&1 == 1 {
			p ^= a
		}

		b >>= 1
		a <<= 1

		if a&order != 0 {
			a ^= r
		}
	}
	return
}

func (r poly) findOthers() {
	// p = 2

	seen := make(map[poly]struct{})

	order := poly(1) << r.degree()
	for a := poly(0); a < order; a++ {
		if _, ok := seen[a]; ok {
			continue
		}
		min := 2 ^ a
		sqr := a
		for {
			seen[sqr] = struct{}{}

			if sqr = r.square(sqr); sqr == a {
				break
			}
			min = r.mul(min, 2^sqr)
		}

		fmt.Printf("%b\n", min)
	}
}
