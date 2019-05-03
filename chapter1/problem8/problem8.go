package main

import "fmt"

func firstRowHasZeros(matrix [][]int) bool {
	for _, v := range matrix[0] {
		if v == 0 {
			return true
		}
	}
	return false
}

func firstColumHasZeros(matrix [][]int) bool {
	for i := range matrix {
		if matrix[i][0] == 0 {
			return true
		}
	}
	return false
}

func nullifyRow(row int, matrix [][]int) {
	for i := range matrix[row] {
		matrix[row][i] = 0
	}
}

func nullifyColumn(column int, matrix [][]int) {
	for i := range matrix {
		matrix[i][column] = 0
	}
}

func nullify(matrix [][]int) {
	_firstRowHasZeros := firstRowHasZeros(matrix)
	_firstColumHasZeros := firstColumHasZeros(matrix)
	for i := range matrix {
		if i == 0 {
			continue
		}
		for j, v := range matrix[i] {
			if j == 0 {
				continue
			}
			if v == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := range matrix {
		if i == 0 {
			continue
		}
		if matrix[i][0] == 0 {
			nullifyRow(i, matrix)
		}
	}

	for i, v := range matrix[0] {
		if i == 0 {
			continue
		}
		if v == 0 {
			nullifyColumn(i, matrix)
		}
	}

	if _firstColumHasZeros {
		nullifyColumn(0, matrix)
	}
	if _firstRowHasZeros {
		nullifyRow(0, matrix)
	}
}

func main() {
	// matrix1 := [][]int{
	// 	{1, 1, 1, 1},
	// 	{1, 0, 1, 1},
	// 	{1, 1, 1, 0},
	// }
	// nullify(matrix1)
	// fmt.Println(matrix1)

	// matrix2 := [][]int{
	// 	{1, 0, 1, 1},
	// 	{1, 0, 1, 1},
	// 	{1, 1, 1, 0},
	// }
	// nullify(matrix2)
	// fmt.Println(matrix2)

	matrix3 := [][]int{
		{1, 0, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 0, 20},
	}
	nullify(matrix3)
	fmt.Println(matrix3)
}
