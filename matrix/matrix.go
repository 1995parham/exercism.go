package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix is a 2D collection of integer.
type Matrix [][]int

// New creates new matrix from given string representation
// of a matrix.
func New(in string) (Matrix, error) {
	m := Matrix(make([][]int, 0))

	// each line represents a row
	rows := strings.Split(in, "\n")

	for _, row := range rows {
		// columns are separated with spaces
		cols := strings.Fields(row)

		// creates new row
		newRow := make([]int, 0)

		for _, col := range cols {
			// convert string cell to number
			n, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			// adds new number into new row
			newRow = append(newRow, n)
		}
		// adds new row into matrix
		m = append(m, newRow)
	}

	// validate that each row size must equal with other rows
	for i := 1; i < len(m); i++ {
		if len(m[i]) != len(m[i-1]) {
			return nil, fmt.Errorf("each row size must equat with other rows size")
		}
	}

	return m, nil
}

// Rows returns rows of matrix.
func (m Matrix) Rows() [][]int {
	rows := make([][]int, 0)

	for _, row := range m {
		newRow := make([]int, len(row))
		copy(newRow, row)
		rows = append(rows, newRow)
	}

	return rows
}

// Cols returns cols of matrix.
func (m Matrix) Cols() [][]int {
	cols := make([][]int, 0)

	// for each column
	for i := 0; i < len(m[0]); i++ {
		newCol := make([]int, 0)
		for _, row := range m {
			newCol = append(newCol, row[i])
		}

		cols = append(cols, newCol)
	}

	return cols
}

// Set changes value in given row and column to new given value.
func (m Matrix) Set(row, col, value int) bool {
	if row < len(m) && row >= 0 {
		if col < len(m[row]) && col >= 0 {
			m[row][col] = value

			return true
		}
	}

	return false
}
