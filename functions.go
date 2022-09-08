package matrix

import (
	"math"
	"math/rand"
)

// Zeros creates a matrix of 0s of the given size
func Zeros(row_count, col_count int) Matrix {
	rows := make([]Row, row_count)
	for i, _ := range rows {
		rows[i] = make(Row, col_count)
	}
	var mx Matrix = rows
	return mx
}

// Generate creates a matrix with all values of the given value.
func Generate(row_count, col_count int, value Col) Matrix {
	rows := make([]Row, row_count)
	for i, _ := range rows {
		rows[i] = make(Row, col_count)
		for j := range rows[i] {
			rows[i][j] = value
		}
	}
	var mx Matrix = rows
	return mx
}

func GenerateRand(row_count, col_count int) Matrix {
	rows := make([]Row, row_count)
	for i, _ := range rows {
		rows[i] = make(Row, col_count)
		for j := range rows[i] {
			value := math.Floor(rand.Float64()*10000) / 100
			rows[i][j] = Col(value)
		}
	}
	var mx Matrix = rows
	return mx
}

// AsMatrix converts an array in the form of [][]float64 to a matrix
func AsMatrix(arr [][]float64) Matrix {
	row_count := len(arr)
	if row_count <= 0 {
		panic("no rows in slice")
	}
	col_count := len(arr[0])
	matrix := Zeros(row_count, col_count)
	for row := range arr {
		for col := range arr[row] {
			matrix[row][col] = Col(arr[row][col])
		}
	}
	return matrix
}
