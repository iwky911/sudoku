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

func TestMatrixCreation(t *testing.T) {
	m := createSparseMatrix(4)
	for i := 0; i < len(m.headers); i++ {
		header := &m.headers[i]
		if header.ncells != 4 {
			t.Errorf("expecting 4 cells per column in column %i", i)
		}
		if header.last == nil {
			t.Errorf("expecting to be able to reach the last")
		}
		c := header.last
		for j := 0; j < header.ncells; j++ {
			c = c.down
		}
		if c != header.last {
			t.Errorf("Doesn't have the right number of actual rows!")
		}
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
