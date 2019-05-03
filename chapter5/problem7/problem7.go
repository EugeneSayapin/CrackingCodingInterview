package main

import (
	"fmt"
)

func swapBits(n uint) uint {
	fmt.Printf("%b\n", n)
	fmt.Printf("%b\n", uint(0x55555555))
	fmt.Printf("%b\n", uint(0xaaaaaaaa))
	fmt.Printf("even: %b\n", (n & 0x55555555))
	fmt.Printf(" odd:  %b\n", (n & 0xaaaaaaaa))
	result := ((n & 0x55555555) << 1) | ((n & 0xaaaaaaaa) >> 1)
	fmt.Printf("%b\n", result)
	return result
}

func main() {
	fmt.Println(swapBits(13948))
}

/*
11011001111100
11100110111100
*/
