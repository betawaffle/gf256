package gf256

import "testing"

func TestString(t *testing.T) {
	a := Element(1<<4 | 1<<3 | 1<<1 | 1<<0)
	if a.String() != "x^4 + x^3 + x + 1" {
		t.Fatalf("unexpected result from Element(%d).String(): %[1]s", a)
	}
}
