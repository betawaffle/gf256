package gf256

import "testing"

func TestAxioms(t *testing.T) {
	// brute-force
	for i := 0; i < 256; i++ {
		a := Element(i)
		testUnaryProperties(t, a)

		for j := 0; j < 256; j++ {
			b := Element(j)
			testBinaryProperties(t, a, b)

			if testing.Short() {
				continue
			}

			for k := 0; k < 256; k++ {
				c := Element(k)
				testTernaryProperties(t, a, b, c)
			}
		}
	}
}

func testUnaryProperties(t *testing.T, a Element) {
	// a + 0 = a
	if Add(a, 0) != a {
		t.Fatalf("invalid additive identity for (%s)", a)
	}

	// a * 1 = a
	if Mul(a, 1) != a {
		t.Fatalf("invalid multiplicative identity for (%s)", a)
	}

	// a * 0 = 0
	if Mul(a, 0) != 0 {
		t.Fatalf("invalid result for multiplication by zero for (%s)", a)
	}

	// a + (-a) = 0
	if Add(a, Neg(a)) != 0 {
		t.Fatalf("invalid additive inverse for (%s)", a)
	}

	// a * a^-1 = 1
	if a != 0 && Mul(a, Inv(a)) != 1 {
		t.Fatalf("invalid multiplicative inverse for (%s)", a)
	}

	// 1 / a = a^-1
	if a != 0 && Div(1, a) != Inv(a) {
		t.Fatalf("division disagrees with multiplicative inverse for (%s)", a)
	}
}

func testBinaryProperties(t *testing.T, a, b Element) {
	// a + b = b + a
	assertCommutative(t, Add, a, b, "addition", "+")

	// a * b = b * a
	assertCommutative(t, Mul, a, b, "multiplication", "*")

	// a - b = a + (-b)
	if Sub(a, b) != Add(a, Neg(b)) {
		t.Fatalf("subtraction disagrees with addition of additive inverse for (%s) and (%s)", a, b)
	}

	// ensure the table-based version matches the algorithmic version
	// if Mul(a, b) != MulBits(a, b) {
	// 	t.Fatalf("multiplication disagrees with bit-by-bit multiplication for (%s) and (%s)", a, b)
	// }
}

func testTernaryProperties(t *testing.T, a, b, c Element) {
	// a + (b + c) = (a + b) + c
	assertAssociative(t, Add, a, b, c, "addition", "+")

	// a * (b * c) = (a * b) * c
	assertAssociative(t, Mul, a, b, c, "multiplication", "*")

	// a * (b + c) = (a * b) + (a * c)
	assertDistributive(t, Mul, Add, a, b, c, "multiplication", "addition")
}

type binaryOp func(x, y Element) Element

func assertAssociative(t *testing.T, op binaryOp, a, b, c Element, name, symbol string) bool {
	if op(a, op(b, c)) != op(op(a, b), c) {
		t.Fatalf("%s not associative for (%[3]s) %[2]s (%[4]s) %[2]s (%[5]s)", name, symbol, a, b, c)
		return false
	}
	return true
}

func assertCommutative(t *testing.T, op binaryOp, a, b Element, name, symbol string) bool {
	if op(a, b) != op(b, a) {
		t.Fatalf("%s not commutative for (%[3]s) %[2]s (%[4]s)", name, symbol, a, b)
		return false
	}
	return true
}

func assertDistributive(t *testing.T, op1, op2 binaryOp, a, b, c Element, name1, name2 string) bool {
	if op1(a, op2(b, c)) != op2(op1(a, b), op1(a, c)) {
		t.Fatalf("%s not distributive over %s for (%s), (%s), and (%s)", name1, name2, a, b, c)
		return false
	}
	return true
}
