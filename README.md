# Matrix Package
The Matrix package allows you to easily perform basic matrix operations in the go programming language.

> **Warning**: Still in development!

## Installation
```
go get github.com/sschrs/matrix
```

## Basic Usage
### Create A Matrix
```go
import "github.com/sschrs/matrix"

func createMatrix() {
custom_matrix := matrix.Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}} // 3x3 Matrix
zero_matrix := matrix.Zeros(3, 3)          // 3x3 Matrix with zeros
one_matrix := matrix.Generate(5, 5, 1)     // 5x5 Matrix with ones
random_matrix := matrix.GenerateRand(4, 4) // 4x4 Matrix with random values
unit_matrix := matrix.UnitMatrix(3,3) // 3x3 Unit Matrix
}
```

### Transpose and Inverse of a Matrix
```go
custom_matrix := matrix.Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
transpose := custom_matrix.T() // returns transpose of a matrix
inverse := custom_matrix.Inv() // returns inverse of a matrix
```

### Matrix Multiplication

```go
matrix := matrix.Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
transpose := custom_matrix.T()
custom_matrix.Dot(transpose) // multiply the matrix with its transpose
```

### Apply

```go
// Multiply all the values in the matrix by 3 and add 10
applied_matrix := custom_matrix.Apply(func(x matrix.Col) matrix.Col {
    return 3*x + 10
})
```

For more information see the docs: https://github.com/sschrs/matrix/wiki