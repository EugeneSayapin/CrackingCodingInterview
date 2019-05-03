package main

import (
	"fmt"
)

func getPermutations(s string) []string {
	return getPermustationsInternal(s, []string{""})
}

func getPermustationsInternal(s string, perms []string) []string {
	if len(s) == 0 {
		return perms
	}

	var newPerms []string
	newPerms = make([]string, 0)

	for _, perm := range perms {
		for i := 0; i <= len(perm); i++ {
			newPerm := perm[0:i] + string(s[0]) + perm[i:]
			newPerms = append(newPerms, newPerm)
		}
	}
	return getPermustationsInternal(s[1:], newPerms)

}

func main() {
	result := getPermutations("abc")
	for _, s := range result {
		fmt.Println(s)
	}
}
