package main

import (
	"fmt"
)

func mult_intenal(big, small uint) uint {
	fmt.Printf("big: %v, small: %v\n", big, small)
	if small == 0 {
		return 0
	}
	if small == 1 {
		return big
	}

	half := mult_intenal(big, small>>1) << 1
	if small&1 == 0 {
		return half
	}
	return half + big
}

func mult(n, m uint) uint {
	if n > m {
		return mult_intenal(n, m)
	}
	return mult_intenal(m, n)
}

func main() {
	fmt.Println(mult(31, 35))
}
