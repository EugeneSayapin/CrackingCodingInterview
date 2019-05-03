package main

import (
	"fmt"
)

func paint(screen [][]rune, r int, c int, newColor rune) {
	if screen[r][c] == newColor {
		return
	}
	oldColor := screen[r][c]

	paintIntenal(screen, r, c, newColor, oldColor)

}

func paintIntenal(screen [][]rune, r int, c int, newColor rune, oldColor rune) {
	if r < 0 || c < 0 || r == len(screen) || c == len(screen[0]) {
		return
	}
	if screen[r][c] != oldColor {
		return
	}

	screen[r][c] = newColor

	paintIntenal(screen, r-1, c, newColor, oldColor)
	paintIntenal(screen, r+1, c, newColor, oldColor)
	paintIntenal(screen, r, c-1, newColor, oldColor)
	paintIntenal(screen, r, c+1, newColor, oldColor)
}

func printScreen(screen [][]rune) {
	for _, row := range screen {
		for _, value := range row {
			fmt.Print(string(value))
		}
		fmt.Println()
	}

}

func main() {
	screen := [][]rune{
		{'b', 'r', 'r', 'b'},
		{'r', 'r', 'r', 'r'},
		{'y', 'y', 'r', 'b'},
	}
	//screen := screenArray[:][:]
	printScreen(screen)
	fmt.Println("---------------------------")
	paint(screen, 2, 0, 'g')
	printScreen(screen)
}
