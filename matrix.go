package matrix

import "fmt"

type Col float64
type Row []Col
type Matrix []Row

func (matrix Matrix) Print() {
	for _, v := range matrix {
		fmt.Println(v)
	}
}

func (matrix Matrix) Copy() Matrix {
	rows := make([]Row, matrix.Shape()["rows"])
	for i := range rows {
		rows[i] = make(Row, matrix.Shape()["cols"])
		for j := range matrix[i] {
			rows[i][j] = matrix[i][j]
		}
	}
	var newMatrix Matrix = rows
	return newMatrix
}

func (matrix Matrix) Shape() map[string]int {
	shape := make(map[string]int)
	shape["rows"] = len(matrix)
	if shape["rows"] == 0 {
		shape["cols"] = 0
	} else {
		shape["cols"] = len(matrix[0])
	}
	return shape
}

func (matrix Matrix) T() Matrix {
	new_matrix := Zeros(matrix.Shape()["cols"], matrix.Shape()["rows"])
	for row_index, _ := range matrix {
		for col_index := range matrix[row_index] {
			new_matrix[col_index][row_index] = matrix[row_index][col_index]
		}
	}
	return new_matrix
}

func (matrix Matrix) Dot(mx Matrix) Matrix {
	if matrix.Shape()["cols"] != mx.Shape()["rows"] {
		panic("must be n = p for matrix multiplication")
	}
	new_matrix := Zeros(matrix.Shape()["rows"], mx.Shape()["cols"])
	for i := 0; i < matrix.Shape()["rows"]; i++ {
		for j := 0; j < mx.Shape()["cols"]; j++ {
			for k := 0; k < matrix.Shape()["cols"]; k++ {
				new_matrix[i][j] += matrix[i][k] * mx[k][j]
			}
		}
	}
	return new_matrix
}

func (matrix Matrix) Add(mx Matrix) Matrix {
	if matrix.Shape()["cols"] != mx.Shape()["cols"] || matrix.Shape()["rows"] != mx.Shape()["rows"] {
		panic("the matrices must be in the same size.")
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			mx[i][j] = mx[i][j] + matrix[i][j]
		}
	}
	return mx
}

func (matrix Matrix) Subtract(mx Matrix) Matrix {
	if matrix.Shape()["cols"] != mx.Shape()["cols"] || matrix.Shape()["rows"] != mx.Shape()["rows"] {
		panic("the matrices must be in the same size.")
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			mx[i][j] = matrix[i][j] - mx[i][j]
		}
	}
	return mx
}

func (matrix Matrix) Plus(value Col) Matrix {
	rows := make([]Row, matrix.Shape()["rows"])
	for i := range rows {
		rows[i] = make(Row, matrix.Shape()["cols"])
		for j := range matrix[i] {
			rows[i][j] = matrix[i][j] + value
		}

	}
	var newMatrix Matrix = rows
	return newMatrix
}

func (matrix Matrix) Minus(value Col) Matrix {
	return matrix.Plus(-value)
}

func (matrix Matrix) Multiply(value Col) Matrix {
	rows := make([]Row, matrix.Shape()["rows"])
	for i := range rows {
		rows[i] = make(Row, matrix.Shape()["cols"])
		for j := range matrix[i] {
			rows[i][j] = matrix[i][j] * value
		}

	}
	var newMatrix Matrix = rows
	return newMatrix
}

func (matrix Matrix) Divide(value Col) Matrix {
	return matrix.Multiply(1 / value)
}

func (matrix Matrix) Apply(f func(x Col) Col) Matrix {
	rows := make([]Row, matrix.Shape()["rows"])
	for i := range rows {
		rows[i] = make(Row, matrix.Shape()["cols"])
		for j := range matrix[i] {
			rows[i][j] = f(matrix[i][j])
		}

	}
	var newMatrix Matrix = rows
	return newMatrix
}

func (matrix Matrix) Sum() float64 {
	var sum Col = 0
	for _, r := range matrix {
		for _, c := range r {
			sum += c
		}
	}
	return float64(sum)
}

func (matrix Matrix) Mean() float64 {
	return matrix.Sum() / float64((matrix.Shape()["cols"] * matrix.Shape()["rows"]))
}

func (matrix Matrix) JoinRow(rows []Row, index int) Matrix {
	var mx Matrix = rows
	if mx.Shape()["cols"] != matrix.Shape()["cols"] {
		panic("the matrices must have same column size.")
	}

	if index > len(matrix)-1 {
		panic("index out of range")
	}

	if index == -1 {
		index = len(matrix)
	}

	newMatrix := Zeros(matrix.Shape()["rows"]+mx.Shape()["rows"], matrix.Shape()["cols"])

	row_index := 0
	for i := 0; i < index; i++ {
		newMatrix[i] = matrix[i]
		row_index++
	}

	for i := index; i < index+len(mx); i++ {
		newMatrix[i] = mx[i-index]
		row_index++
	}

	for i := index; i < len(matrix); i++ {
		newMatrix[row_index] = matrix[i]
		row_index++
	}

	return newMatrix
}

func (matrix Matrix) MultiplyRow(rowIndex int, value Col) Matrix {
	if rowIndex >= len(matrix) {
		panic("index out of range")
	}
	newMatrix := matrix.Copy()
	for i := range newMatrix[rowIndex] {
		newMatrix[rowIndex][i] *= value
	}
	return newMatrix
}

func (matrix Matrix) SwapRows(firstIndex, secondIndex int) Matrix {
	if firstIndex >= len(matrix) || secondIndex >= len(matrix) {
		panic("index out of range")
	}
	newMatrix := matrix.Copy()
	firstRow := newMatrix[firstIndex]
	newMatrix[firstIndex] = newMatrix[secondIndex]
	newMatrix[secondIndex] = firstRow
	return newMatrix
}

func (matrix Matrix) AddRows(destination, source int) Matrix {
	if destination >= len(matrix) || source >= len(matrix) {
		panic("index out of range")
	}
	newMatrix := matrix.Copy()
	for i := range matrix[destination] {
		newMatrix[source][i] += newMatrix[destination][i]
	}
	return newMatrix
}

func (matrix Matrix) PlusRow(index int, row []Col) Matrix {
	if index >= len(matrix) {
		panic("index out of range")
	}
	newMatrix := matrix.Copy()
	for i := range newMatrix[index] {
		newMatrix[index][i] += row[i]
	}
	return newMatrix
}

func (matrix Matrix) GetColumn(colIndex int) []Col {
	var col []Col
	for i := range matrix {
		col = append(col, matrix[i][colIndex])
	}
	return col
}

func (matrix Matrix) UpperTriangle() Matrix {
	newMatrix := matrix.Copy()

	row_index := len(matrix) - 1
	for newMatrix[0][0] == 0 {
		newMatrix = newMatrix.SwapRows(0, row_index)
		row_index--
		if row_index < 0 {
			panic("all columns is 0 in 0th index")
		}
	}

	for i := range newMatrix {
		multiplication_value := newMatrix[i][i]
		for j := i + 1; j < len(newMatrix); j++ {
			newMatrix = newMatrix.PlusRow(j, newMatrix.MultiplyRow(i, -(newMatrix[j][i] / multiplication_value))[i])
		}
	}
	return newMatrix
}
