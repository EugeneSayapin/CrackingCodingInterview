package main

import "fmt"

func IsUnique(input string) bool{
	runes := []rune(input)
	runes = qsort(runes);
	for i:=1; i < len(runes); i++ {
		if runes[i-1] == runes[i] {
			return false;
		}
	}
	return true;
}
func qsort(a []rune) []rune {
	if len(a) < 2 { return a }
  
	left, right := 0, len(a) - 1
  
	// Pick a pivot
	pivotIndex := left
  
	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
  
	// Pile elements smaller than the pivot on the left
	for i := range a {
	  if a[i] < a[right] {
		a[i], a[left] = a[left], a[i]
		left++
	  }
	}
  
	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]
  
	// Go down the rabbit hole
	qsort(a[:left])
	qsort(a[left + 1:])
  
  
	return a
  }
func main() {
	fmt.Printf("abc is unique ")
	fmt.Println(IsUnique("abc"))
	fmt.Printf("bbc is unique ")
	fmt.Println(IsUnique("bbc"))
	fmt.Printf("Hello world is unique ")
	fmt.Println(IsUnique("Hello world"))
	fmt.Printf("Helo wrd is unique ")
	fmt.Println(IsUnique("Helo wrd"))
}