/*
Package gf256 defines operations on elements of the non-prime finite field
GF(2^8).

Elements of GF(2^n) are actually polynomials (with degree < n) of elements in
GF(2).

The machine code generated for the non-table-based functions (Add, Neg, Sub,
Mul, Inv, Div) is expected to be constant-time, but that has not been
rigorously validated.
*/
package gf256

const (
	// Characteristic is the number of times you must add an element to itself to
	// get the additive identity.
	Characteristic = 2

	// Dimension is the number of terms in the polynomial described by each
	// element.
	Dimension = 8

	// Order is the number of elements in the field. It is equal to the
	// characteristic raised to the dimension-th power.
	Order = 1 << Dimension

	// MaxDegree is the highest degree of any element in the field.
	MaxDegree = Dimension - 1

	// MaxElement is the highest element in the field.
	MaxElement = Order - 1
)

const (
	// r represents x^8 + x^4 + x^3 + x + 1 as an element of GF(256).
	//
	// Solving for x, we find:
	//   x^8 + x^4 + x^3 + x + 1 = 0
	//   x^8                     = Neg(x^4 + x^3 + x + 1)
	//   x^8                     =     x^4 + x^3 + x + 1
	r = Element(1<<4 | 1<<3 | 1<<1 | 1<<0)

	// g represents x + 1. It is the smallest generator in GF(256).
	g = Element(1<<1 | 1<<0)
)

// Element represents an element of GF(2^8).
type Element uint8

// Add returns the sum of a and b in GF(256). The implementation is expected to
// be constant-time.
func Add(a, b Element) Element {
	return a ^ b
}

// Neg returns the negation of a in GF(256). The implementation is expected to
// be constant-time.
func Neg(a Element) Element {
	return a // Sub(0, a)
}

// Sub returns the difference of a and b in GF(256). The implementation is
// expected to be constant-time.
func Sub(a, b Element) Element {
	return a ^ b // Add(a, Neg(b))
}

// Mul returns a multiplied by b in GF(256). The implementation is expected to
// be constant-time.
func Mul(a, b Element) Element {
	// a7 = bit 7 from a
	// b0 = bit 0 from b
	var p Element
	for i := 0; i < Dimension; i++ {
		ab0 := a &^ (b&1 - 1)  // ab0 = Mul(a, b0)
		ra7 := r &^ (a>>7 - 1) // ra7 = Mul(R, a7)

		// p = Add(p, ab0)
		// a = Sub(Mul(a, x), ra7)
		// b = Div(b, x)
		p, a, b = p^ab0, a<<1^ra7, b>>1
	}
	return p
}

// Inv returns the multiplicative inverse of a (1 / a) in GF(256). The
// implementation is expected to be constant-time.
func Inv(a Element) Element {
	if a == 0 {
		panic("inverse of Element(0) is undefined")
	}
	var (
		p = Element(1)
		s = a
	)
	// 2 + 4 + 8 + 16 + 32 + 64 + 128 = 254
	for i := 2; i < Order; i <<= 1 {
		s = Mul(s, s)
		p = Mul(p, s)
	}
	return p
}

// Div returns the quotient of a and b in GF(256). The implementation is
// expected to be constant-time.
func Div(a, b Element) Element {
	return Mul(a, Inv(b))
}
