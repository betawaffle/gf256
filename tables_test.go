package gf256

import "testing"

func TestTables(t *testing.T) {
	// brute-force
	for i := 0; i < 256; i++ {
		a := Element(i)

		if a != 0 {
			testInvTable(t, a)
		}

		for j := 0; j < 256; j++ {
			b := Element(j)
			testMulTable(t, a, b)

			if b != 0 {
				testDivTable(t, a, b)
			}
		}
	}
}

func testInvTable(t *testing.T, a Element) {
	expect := Inv(a)
	actual := InvTable(a)

	if expect != actual {
		t.Errorf("expected InvTable(%s) to return %s, got %s", a, expect, actual)
	}
}

func testMulTable(t *testing.T, a, b Element) {
	expect := Mul(a, b)
	actual := MulTable(a, b)

	if expect != actual {
		t.Errorf("expected MulTable(%s, %s) to return %s, got %s", a, b, expect, actual)
	}
}

func testDivTable(t *testing.T, a, b Element) {
	expect := Div(a, b)
	actual := DivTable(a, b)

	if expect != actual {
		t.Errorf("expected DivTable(%s, %s) to return %s, got %s", a, b, expect, actual)
	}
}
