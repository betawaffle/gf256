package gf256

var (
	invTable [256]Element
	expTable [256]Element
	logTable [256]Element
)

func init() {
	invTable[0] = 0
	logTable[0] = 0

	for c, a := Element(0), Element(1); c < 255; c++ {
		expTable[c] = a
		logTable[a] = c

		a = Mul(a, g)
	}
	expTable[255] = expTable[0]

	for i := 1; i < 256; i++ {
		a := Element(i)
		invTable[a] = expTable[255-logTable[a]]
	}
}

func (e Element) eq0() Element {
	e = ^e
	e &= e >> 4
	e &= e >> 2
	e &= e >> 1
	return e
}

// MulTable uses table lookups to multiply elements. The implementation is not
// necessarily constant-time.
func MulTable(a, b Element) Element {
	z := a.eq0() | b.eq0()
	s := uint16(logTable[a]) + uint16(logTable[b])

	// clear all bits if a or b are zero
	return expTable[s%255] & (z - 1)
}

// InvTable uses a table lookup to find the multiplicative inverse of an
// element. The implementation is not necessarily constant-time.
func InvTable(a Element) Element {
	if a == 0 {
		panic("undefined")
	}
	return invTable[a]
}

// DivTable uses table lookups to divide elements. The implementation is not
// constant-time.
func DivTable(a, b Element) Element {
	if b == 0 {
		panic("undefined")
	}
	if a == 0 {
		return 0
	}
	d := int16(logTable[a]) - int16(logTable[b])
	d += (d - 1) >> 15 & 255
	return expTable[d]
}
