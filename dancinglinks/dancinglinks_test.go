package main

import (
	"testing"
)

func TestGetSquare(t *testing.T) {
	if getSquare(2, 3, 9) != 1 {
		t.Errorf("Expecting square 1")
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
