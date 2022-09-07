package matrix

func Zeros(row_count, col_count int) Matrix {
	rows := make([]Row, row_count)
	for i, _ := range rows {
		rows[i] = make(Row, col_count)
	}
	var mx Matrix = rows
	return mx
}

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
