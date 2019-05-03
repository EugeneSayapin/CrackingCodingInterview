package main

import (
	"fmt"
)

func isReplace(s1, s2 string) bool {
	rune1 := []rune(s1)
	rune2 := []rune(s2)
	seenReplace := false
	for i := range rune1 {
		if rune1[i] != rune2[i] {
			if seenReplace {
				return false
			}
			seenReplace = true
		}
	}
	return seenReplace
}
func isInsert(longer, shorter string) bool {
	longerRune := []rune(longer)
	shorterRune := []rune(shorter)
	seenInsert := false
	for longerIndex := range longerRune {
		if longerIndex == len(shorterRune) && !seenInsert {
			return true
		}

		shorterIndex := longerIndex
		if seenInsert {
			shorterIndex--
		}
		if shorterRune[shorterIndex] != longerRune[longerIndex] {
			if seenInsert {
				return false
			}
			seenInsert = true
		}
	}
	return seenInsert
}

func isEdit(s1 string, s2 string) bool {
	if len(s1) == len(s2) {
		return isReplace(s1, s2)
	}
	var longer, shorter string
	if len(s1) > len(s2) {
		longer = s1
		shorter = s2
	} else {
		longer = s2
		shorter = s1
	}

	if len(longer)-len(shorter) != 1 {
		return false
	}
	return isInsert(longer, shorter)

}

func main() {
	fmt.Printf("%v, %v -> %v\n", "pale", "ple", isEdit("pale", "ple"))
	fmt.Printf("%v, %v -> %v\n", "pales", "pale", isEdit("pales", "pale"))
	fmt.Printf("%v, %v -> %v\n", "pale", "bale", isEdit("pale", "bale"))
	fmt.Printf("%v, %v -> %v\n", "pale", "bake", isEdit("pale", "bake"))

}
