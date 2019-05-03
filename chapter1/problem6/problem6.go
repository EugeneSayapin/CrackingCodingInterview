package main

import (
	"bytes"
	"fmt"
)

func compress(input string) string {
	if input == "" {
		return ""
	}
	runes := []rune(input)
	resultSymbols := make([]rune, len(runes))
	resultCounts := make([]int, len(runes))
	resultsIndex := 0
	resultSymbols[0] = runes[0]
	resultCounts[0] = 1
	for i := 1; i < len(runes); i++ {
		if runes[i] == resultSymbols[resultsIndex] {
			resultCounts[resultsIndex]++
		} else {
			resultsIndex++
			resultSymbols[resultsIndex] = runes[i]
			resultCounts[resultsIndex] = 1
		}
	}

	var buffer bytes.Buffer
	for i := 0; i <= resultsIndex; i++ {
		buffer.WriteRune(resultSymbols[i])
		buffer.WriteString(fmt.Sprintf("%d", resultCounts[i]))
	}
	var result = buffer.String()

	if len(result) < len(input) {
		return result
	} else {
		return input
	}

}

func main() {
	fmt.Println(compress("aabcccccaaa"))
	fmt.Println(compress("aacc"))
}
