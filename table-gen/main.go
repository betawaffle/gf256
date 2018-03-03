package main

import "fmt"

// 2^8  - 1 = 3 * 5 * 17
// 2^16 - 1 = 3 * 5 * 17 * 257
// 2^32 - 1 = 3 * 5 * 17 * 257 * 65537
var factors = [...]uint{3, 5, 17, 257, 65537}

func main() {
	for i, r := range rs {
		fmt.Printf("generator for r%-2d = %08b: %d\n", i+1, r, gs[i])
	}
}
