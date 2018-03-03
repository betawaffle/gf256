package gf256

import (
	"testing"
)

func TestMul(t *testing.T) {
	x := Element(1 << 1)
	y := x

	for i := 2; i < 256; i++ {
		y = Mul(y, x)
	}
	if y != 1 {
		t.Fatalf("x^255 != 1")
	}

	if Mul(1<<7, 2) != r {
		t.Fatalf("x^7 * x != %s", r)
	}
}

func BenchmarkMul(b *testing.B) {
	x := Element(2)
	y := Element(3)

	for i := 0; i < b.N; i++ {
		x = Mul(x, y)
	}
	if x == 0 {
		b.Fatalf("x == 0")
	}
}

func BenchmarkInv(b *testing.B) {
	x := Element(2)

	for i := 0; i < b.N; i++ {
		x = Inv(x)
	}
	if x == 0 {
		b.Fatalf("x == 0")
	}
}

func BenchmarkDiv(b *testing.B) {
	x := Element(2)
	y := Element(3)

	for i := 0; i < b.N; i++ {
		x = Div(x, y)
	}
	if x == 0 {
		b.Fatalf("x == 0")
	}
}
