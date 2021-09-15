package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

// New builds a matrix from a valid string
func New(s string) (Matrix, error) {
	// split string into rows
	lines := strings.Split(s, "\n")
	rows := make([][]int, 0)
	for _, l := range lines {
		// clean rows from white spaces
		trimmed := strings.TrimSpace(l)
		// split row into cells
		cells := strings.Split(trimmed, " ")
		numCells := make([]int, 0)
		for _, c := range cells {
			// convert cell to integer
			num, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			// build a row
			numCells = append(numCells, num)
		}
		// build the matrix
		rows = append(rows, numCells)
	}
	// validate amount of columns (same for each row)
	for _, h := range rows {
		if len(h) != len(rows[0]) {
			return nil, errors.New("invalid matrix")
		}
	}
	return rows, nil
}

// Cols returns a copy of each column from the matrix
func (m Matrix) Cols() [][]int {
	result := make([][]int, 0)
	for i := 0; i < len(m[0]); i++ {
		col := make([]int, 0)
		for j := 0; j < len(m); j++ {
			col = append(col, m[j][i])
		}
		result = append(result, col)
	}
	return result
}

// Rows returns a copy of each row from the matrix
func (m Matrix) Rows() [][]int {
	result := make([][]int, 0)
	for _, l := range m {
		r := make([]int, 0)
		for _, k := range l {
			r = append(r, k)
		}
		result = append(result, r)
	}
	return result
}

// Set allows to modify a value from a specific cell
func (m Matrix) Set(row, column, value int) bool {
	if row > len(m)-1 || row < 0 || column > len(m[0])-1 || column < 0 {
		return false // out of matrix
	}
	m[row][column] = value
	return true
}
