package main

import (
	"fmt"
)

func getPermutations(s string) []string {
	dict := getCountDict(s)
	result := []string{}
	getPermutationsInternal(dict, len(s), "", &result)
	return result
}

func getPermutationsInternal(dict map[rune]int, remaning int, prefix string, perms *[]string) {
	if remaning == 0 {
		*perms = append(*perms, prefix)
	}

	for char, count := range dict {
		if count == 0 {
			continue
		}
		dict[char] = count - 1
		getPermutationsInternal(dict, remaning-1, prefix+string(char), perms)
		dict[char] = count

	}

}

func getCountDict(s string) map[rune]int {
	runes := []rune(s)
	var result map[rune]int
	result = make(map[rune]int)
	for _, rune := range runes {
		count := result[rune]
		if result == nil {
			result[rune] = 0
		}
		result[rune] = count + 1
	}
	return result
}

func main() {
	perms := getPermutations("abccd")
	for _, s := range perms {
		fmt.Println(s)
	}
}
