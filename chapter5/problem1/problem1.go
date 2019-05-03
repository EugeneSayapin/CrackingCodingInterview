package main

import (
	"fmt"
	"strconv"
)

func insert(m, n, i, j uint) uint {
	ones := uint(((1 << (j - i + 1)) - 1) << i)
	fmt.Printf("%b\n", ones)
	n = n & ^ones

	return n | (m << i)

}

func main() {
	n, _ := strconv.ParseInt("10000111100", 2, 32)
	m, _ := strconv.ParseInt("10011", 2, 32)
	fmt.Printf("%b\n", n)
	fmt.Printf("%b\n", insert(uint(m), uint(n), uint(2), uint(6)))
}
