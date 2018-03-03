package main

type tab8 struct {
	log [256]poly8
	exp [256]poly8
}

func (t *tab8) gen(r, g poly8) {
	t.log[0] = 0

	for i, a := poly8(0), poly8(1); i < 255; i, a = i+1, r.mul(a, g) {
		t.exp[i] = a
		t.log[a] = i
	}

	t.exp[255] = t.exp[0]
}

func (t *tab8) mul(a, b poly8) poly8 {
	z := a.eq0() | b.eq0()
	s := uint16(t.log[a]) + uint16(t.log[b])

	// clear all bits if a or b are zero
	return t.exp[s%255] & (z - 1)
}

func (t *tab8) mulGroupOrder(a poly8) int {
	for k, b := 1, a; k < 256; k++ {
		if b = t.mul(b, a); b == a {
			return k
		}
	}
	return 0
}

func (t *tab8) inv(a poly8) poly8 {
	if a == 0 {
		panic("undefined")
	}
	return t.div(1, a)
}

func (t *tab8) div(a, b poly8) poly8 {
	if b == 0 {
		panic("undefined")
	}
	d := int16(t.log[a]) - int16(t.log[b])
	d += (d - 1) >> 15 & 255
	return t.exp[d]
}
