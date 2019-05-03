package main

import (
	"fmt"
	"strings"
)
func isPalindPerm(s string)bool {
	symbols := make(map[rune]int);
	runes := []rune(strings.ToLower(s));
	for _, char := range runes {
		if char == ' ' {
			continue
		}
		symbols[char]++;
	}
	numberOfOdds := 0;
	for _, count := range symbols{
		if count % 2 == 1 {
			numberOfOdds++
		}
	}
	return numberOfOdds < 2;

}

func main() {
	fmt.Println(isPalindPerm("Tact Coa"))
}