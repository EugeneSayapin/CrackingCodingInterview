package main

import (
	"fmt"
)

func generateParens(n int) []string {
	var result []string
	result = make([]string, 0)
	generateParensIntenal(n, n, &result, "")
	return result
}

func generateParensIntenal(leftRem int, rightRem int, list *[]string, current string) {
	if leftRem == 0 && rightRem == 0 {
		*list = append(*list, current)
	}

	if leftRem > 0 {
		generateParensIntenal(leftRem-1, rightRem, list, current+"(")
	}

	if rightRem > leftRem {
		generateParensIntenal(leftRem, rightRem-1, list, current+")")
	}

}

func main() {

	toPrint := generateParens(4)

	for _, s := range toPrint {
		fmt.Println(s)
	}

}
