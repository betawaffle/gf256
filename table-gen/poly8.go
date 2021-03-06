package main

const (
	r1  = poly8(1<<4 | 1<<3 | 1<<1 | 1<<0)               // x^8 + x^4 + x^3 + x + 1
	r2  = poly8(1<<4 | 1<<3 | 1<<2 | 1<<0)               // x^8 + x^4 + x^3 + x^2 + 1
	r3  = poly8(1<<5 | 1<<3 | 1<<1 | 1<<0)               // x^8 + x^5 + x^3 + x + 1
	r4  = poly8(1<<5 | 1<<3 | 1<<2 | 1<<0)               // x^8 + x^5 + x^3 + x^2 + 1
	r5  = poly8(1<<5 | 1<<4 | 1<<3 | 1<<0)               // x^8 + x^5 + x^4 + x^3 + 1
	r6  = poly8(1<<5 | 1<<4 | 1<<3 | 1<<2 | 1<<1 | 1<<0) // x^8 + x^5 + x^4 + x^3 + x^2 + x + 1
	r7  = poly8(1<<6 | 1<<3 | 1<<2 | 1<<0)               // x^8 + x^6 + x^3 + x^2 + 1
	r8  = poly8(1<<6 | 1<<4 | 1<<3 | 1<<2 | 1<<1 | 1<<0) // x^8 + x^6 + x^4 + x^3 + x^2 + x + 1
	r9  = poly8(1<<6 | 1<<5 | 1<<1 | 1<<0)               // x^8 + x^6 + x^5 + x + 1
	r10 = poly8(1<<6 | 1<<5 | 1<<2 | 1<<0)               // x^8 + x^6 + x^5 + x^2 + 1
	r11 = poly8(1<<6 | 1<<5 | 1<<3 | 1<<0)               // x^8 + x^6 + x^5 + x^3 + 1
	r12 = poly8(1<<6 | 1<<5 | 1<<4 | 1<<0)               // x^8 + x^6 + x^5 + x^4 + 1
	r13 = poly8(1<<6 | 1<<5 | 1<<4 | 1<<2 | 1<<1 | 1<<0) // x^8 + x^6 + x^5 + x^4 + x^2 + x + 1
	r14 = poly8(1<<6 | 1<<5 | 1<<4 | 1<<3 | 1<<1 | 1<<0) // x^8 + x^6 + x^5 + x^4 + x^3 + x + 1
	r15 = poly8(1<<7 | 1<<2 | 1<<1 | 1<<0)               // x^8 + x^7 + x^2 + x + 1
	r16 = poly8(1<<7 | 1<<3 | 1<<1 | 1<<0)               // x^8 + x^7 + x^3 + x + 1
	r17 = poly8(1<<7 | 1<<3 | 1<<2 | 1<<0)               // x^8 + x^7 + x^3 + x^2 + 1
	r18 = poly8(1<<7 | 1<<4 | 1<<3 | 1<<2 | 1<<1 | 1<<0) // x^8 + x^7 + x^4 + x^3 + x^2 + x + 1
	r19 = poly8(1<<7 | 1<<5 | 1<<1 | 1<<0)               // x^8 + x^7 + x^5 + x + 1
	r20 = poly8(1<<7 | 1<<5 | 1<<3 | 1<<0)               // x^8 + x^7 + x^5 + x^3 + 1
	r21 = poly8(1<<7 | 1<<5 | 1<<4 | 1<<0)               // x^8 + x^7 + x^5 + x^4 + 1
	r22 = poly8(1<<7 | 1<<5 | 1<<4 | 1<<3 | 1<<2 | 1<<0) // x^8 + x^7 + x^5 + x^4 + x^3 + x^2 + 1
	r23 = poly8(1<<7 | 1<<6 | 1<<1 | 1<<0)               // x^8 + x^7 + x^6 + x + 1
	r24 = poly8(1<<7 | 1<<6 | 1<<3 | 1<<2 | 1<<1 | 1<<0) // x^8 + x^7 + x^6 + x^3 + x^2 + x + 1
	r25 = poly8(1<<7 | 1<<6 | 1<<4 | 1<<2 | 1<<1 | 1<<0) // x^8 + x^7 + x^6 + x^4 + x^2 + x + 1
	r26 = poly8(1<<7 | 1<<6 | 1<<4 | 1<<3 | 1<<2 | 1<<0) // x^8 + x^7 + x^6 + x^4 + x^3 + x^2 + 1
	r27 = poly8(1<<7 | 1<<6 | 1<<5 | 1<<2 | 1<<1 | 1<<0) // x^8 + x^7 + x^6 + x^5 + x^2 + x + 1
	r28 = poly8(1<<7 | 1<<6 | 1<<5 | 1<<4 | 1<<1 | 1<<0) // x^8 + x^7 + x^6 + x^5 + x^4 + x + 1
	r29 = poly8(1<<7 | 1<<6 | 1<<5 | 1<<4 | 1<<2 | 1<<0) // x^8 + x^7 + x^6 + x^5 + x^4 + x^2 + 1
	r30 = poly8(1<<7 | 1<<6 | 1<<5 | 1<<4 | 1<<3 | 1<<0) // x^8 + x^7 + x^6 + x^5 + x^4 + x^3 + 1
)

var rs = [30]poly8{
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30,
}

var gs = [30]poly8{}

func init() {
	for i, n := 2, 0; i < 256; i++ {
		a := poly8(i)

		for j, r := range rs {
			if gs[j] != 0 {
				continue
			}
			if r.mulGroupOrder(a) == 255 {
				gs[j] = a
				n++
			}
			if n == len(gs) {
				return
			}
		}
	}
}

type poly8 uint8

func (r poly8) eq0() poly8 {
	r = ^r
	r &= r >> 4
	r &= r >> 2
	r &= r >> 1
	return r
}

func (r poly8) deg() (n uint) {
	for r > 1 {
		r >>= 1
		n++
	}
	return
}

func (r poly8) mul(a, b poly8) poly8 {
	var p poly8
	for i := 0; i < 8; i++ {
		ab0 := a &^ (b&1 - 1)
		ra7 := r &^ (a>>7 - 1)

		p, a, b = p^ab0, a<<1^ra7, b>>1
	}
	return p
}

func (r poly8) mulGroupOrder(a poly8) int {
	for k, b := 1, a; k < 256; k++ {
		if b = r.mul(b, a); b == a {
			return k
		}
	}
	return 0
}

func (r poly8) inv(a poly8) poly8 {
	if a == 0 {
		panic("undefined")
	}
	var (
		p = poly8(1)
		s = a
	)
	// 2 + 4 + 8 + 16 + 32 + 64 + 128 = 254
	for i := 2; i < 256; i <<= 1 {
		s = r.mul(s, s)
		p = r.mul(p, s)
	}
	return p
}
