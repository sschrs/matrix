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

// GenerateRow creates a row with all values of the given value and returns it as a Row type
func GenerateRow(row_size int, value float64) Row {
	row := make(Row, row_size)
	for i := range row {
		row[i] = Col(value)
	}
	return row
}

// GenerateColumn creates a column with all values of the given value as a []Col type
func GenerateColumn(col_size int, value float64) []Col {
	col := make([]Col, col_size)
	for i := range col {
		col[i] = Col(value)
	}
	return col
}

// UnitMatrix creates a unit matrix and returns it
func UnitMatrix(row_count, col_count int) Matrix {
	rows := make([]Row, row_count)
	for i, _ := range rows {
		rows[i] = make(Row, col_count)
		for j := range rows[i] {
			if i == j {
				rows[i][j] = 1
			} else {
				rows[i][j] = 0
			}
		}
	}
	var mx Matrix = rows
	return mx
}
