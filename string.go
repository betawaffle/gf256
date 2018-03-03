package gf256

var termStrings = [...]string{"1", "x", "x^2", "x^3", "x^4", "x^5", "x^6", "x^7"}

func (e Element) String() string {
	if e == 0 {
		return "0"
	}
	var s string
	for i := uint(0); i < Dimension; i++ {
		if e>>i&1 == 1 {
			if s == "" {
				s = termStrings[i]
			} else {
				s = termStrings[i] + " + " + s
			}
		}
	}
	return s
}
