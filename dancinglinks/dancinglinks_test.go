package main

import (
	"testing"
)

func TestGetSquare(t *testing.T) {
	// row 2, column 3, size 9 (sqrt = 3)
	if getSquare(2, 3, 3) != 1 {
		t.Errorf("Expecting square 1")
	}

	// row 5, column 5, size 9 (sqrt = 3)
	result := getSquare(5, 5, 3)
	if result != 4 {
		t.Errorf("Expecting square 4 but got %i", result)
	}
}

func TestAffectationToCode(t *testing.T) {
	cases := []struct {
		row, column, value, size, code int
	}{
		{0, 0, 0, 4, 0},
		{1, 3, 2, 4, 30},
	}

	for _, tc := range cases {
		result := AffectationToCode(tc.row, tc.column, tc.value, tc.size)
		if result != tc.code {
			t.Errorf("Expected %i but got %i", tc.code, result)
		}

		r, c, v := CodeToAffectation(tc.code, tc.size)
		if r != tc.row || c != tc.column || v != tc.value {
			t.Errorf("Wrong value for code %i", tc.code)
		}
	}
}

func TestMatrixCreation(t *testing.T) {
	m := createSparseMatrix(4)
	if !isSparseMatrixCorrect(m, t) {
		t.Errorf("Matrix created wasn't correct")
	}
}

func TestRemoveColumn(t *testing.T) {
	m := createSparseMatrix(4)

	m.headers[5].RemoveColumn()
	m.headers[5].AddColumn()
	if !isSparseMatrixCorrect(m, t) {
		t.Errorf("Matrix created wasn't correct")
	}
}

func TestSelectValue(t *testing.T) {
	m := createSparseMatrix(4)

	cell := m.headers[16].last.top
	cell.RemoveAllAffectedColumns()
	cell.AddAllAffectedColumns()

	if !isSparseMatrixCorrect(m, t) {
		t.Errorf("Matrix created wasn't correct")
	}
}
func TestSimpleSudokuSolvable(t *testing.T) {
	m := createSparseMatrix(4)
	if !m.Solvable() {
		t.Errorf("matrix should be solvable")
	}
}

func TestGettingSmallestColumn(t *testing.T) {
	m := createSparseMatrix(4)

	expected := &m.headers[3]
	expected.ncells = expected.ncells - 1
	found := getSmallestColumn(m.head)
	if found != expected {
		t.Errorf("smallest column wasn't the expected one")
	}
}

func isSparseMatrixCorrect(m *SparseMatrix, t *testing.T) bool {
	success := true
	for i := 0; i < len(m.headers); i++ {
		header := &m.headers[i]
		if header.ncells != m.size {
			t.Errorf("expecting %i cells per column in column %i", m.size, i)
			success = false
		}
		if header.last == nil {
			t.Errorf("expecting to be able to reach the last")
			success = false
		}
		c := header.last
		for j := 0; j < header.ncells; j++ {
			c = c.down
		}
		if c != header.last {
			t.Errorf("Doesn't have the right number of actual rows!")
			success = false
		}
	}
	return success
}
