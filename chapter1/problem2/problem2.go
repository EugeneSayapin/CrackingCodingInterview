package main

import "fmt"

func isPermutation(s1 string, s2 string) bool {
	m := make(map[rune]int);
	a1 := []rune(s1);
	a2 := []rune(s2);
	for _, c := range a1 {
		m[c]++;
	}
	for _, c := range a2 {
		m[c]--;
	}

	for _, count := range m {
		if count != 0 {
			return false;
		}
	}
	return true;

}

func main() {
	var s1 = "carrot";
	var s2 = "torrrac";
	fmt.Printf(s1);
	fmt.Printf( " is a permutation of " );
	fmt.Printf(s1);
	fmt.Printf(": ");
	fmt.Println(isPermutation(s1, s2));
}