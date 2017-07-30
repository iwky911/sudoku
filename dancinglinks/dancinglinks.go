package main

import (
	"fmt"
	"math"
)

/*
Contains the struct
* first n elt: validate row i contains v
* second n elts: validates column i contains v
* final n elts: validates block i contains v
*/

var originalMatrix [][]int

type SparseMatrix struct {
	headers   []Header
	head      *Header
	size      int
	iteration int
}

type Cell struct {
	top, down, left, right, head *Cell
	header                       *Header
	code                         int
}

func newCell(code int) *Cell {
	cell := new(Cell)
	cell.code = code
	return cell
}

type Header struct {
	left, right *Header
	last        *Cell
	ncells      int
	m           *SparseMatrix
}

func addCellToColumn(sparse *SparseMatrix, index int, cell *Cell) {
	var header *Header
	header = &sparse.headers[index]
	header.ncells++
	cell.header = header
	if header.last == nil {
		header.last = cell
		cell.top = cell
		cell.down = cell
		return
	}

	cell.top = header.last
	cell.down = header.last.down
	cell.top.down = cell
	cell.down.top = cell
}

func createSparseMatrix(size int) *SparseMatrix {
	// Column (constraints) are ordered in the following way:
	// - N * 9 + X (where N < 9 and X < 9): Value X is in column N
	// - 9*9 + N * 9 + X (where N < 9 and X < 9): Value X is in row N
	// - 2*9*9 + N * 9 + X (where N < 9 and X < 9): Value X is in square N
	// - 3*9*9 + R * 9 + C (where C is the colum and R is the row): cell in row R and column C has a value.
	var nColumn = size * size * 4
	sizesqrt := int(math.Sqrt(float64(size)))
	sparse := &SparseMatrix{make([]Header, nColumn), nil, size, 0}
	for h := 0; h < nColumn; h++ {
		var prev = (h - 1 + nColumn) % nColumn
		var next = (h + 1) % nColumn
		sparse.headers[h].right = &sparse.headers[next]
		sparse.headers[h].left = &sparse.headers[prev]
		sparse.headers[h].m = sparse
	}

	sparse.head = &sparse.headers[0]
	for row := 0; row < size; row++ {
		for column := 0; column < size; column++ {
			for value := 0; value < size; value++ {
				// This configuration matches three constraints.
				code := AffectationToCode(row, column, value, size)
				cconstr := newCell(code)
				rconstr := newCell(code)
				sconstr := newCell(code)
				vconstr := newCell(code)

				// wire them horizontally
				cconstr.left = vconstr
				cconstr.right = rconstr
				rconstr.left = cconstr
				rconstr.right = sconstr
				sconstr.left = rconstr
				sconstr.right = vconstr
				vconstr.left = sconstr
				vconstr.right = cconstr

				// Now wire them vertically.
				addCellToColumn(sparse, column*size+value, cconstr)
				addCellToColumn(sparse, size*size+row*size+value, rconstr)
				addCellToColumn(sparse, size*size*2+getSquare(row, column, sizesqrt)*size+value, sconstr)
				addCellToColumn(sparse, getValueConstraintIndex(row, column, size), vconstr)
			}
		}
	}

	fmt.Println("Sparse matrix created")
	return sparse
}

func getValueConstraintIndex(row, column, size int) int {
	return size*size*3 + row*size + column
}

func AffectationToCode(row, column, value, size int) int {
	return value + column*size + row*size*size
}

func CodeToAffectation(code, size int) (int, int, int) {
	value := code % size
	code = code / size
	column := code % size
	code = code / size
	row := code % size
	return row, column, value
}

func getSquare(row, column, sizesqrt int) int {
	return (row/sizesqrt)*sizesqrt + column/sizesqrt
}

func getSmallestColumn(head *Header) *Header {
	if head == nil {
		return nil
	}

	rec := head
	for c := head.right; c != head; c = c.right {
		if c.ncells < rec.ncells {
			rec = c
		}
	}
	return rec
}

// Remove all columns that have a cell on this row, from left to right.
func (cell *Cell) RemoveAllAffectedColumns() {
	cell.header.RemoveColumn()
	for c := cell.right; c != cell; c = c.right {
		c.header.RemoveColumn()
	}
}

// Add all columns that have a cell on this row, from right to left.
func (cell *Cell) AddAllAffectedColumns() {
	for c := cell.left; c != cell; c = c.left {
		c.header.AddColumn()
	}
	cell.header.AddColumn()
}

// Remove this column and all rows it uses.
func (header *Header) RemoveColumn() {
	header.left.right = header.right
	header.right.left = header.left

	if header.left == header {
		header.m.head = nil
	} else {
		header.m.head = header.right
	}

	// Remove the rows going down.
	cell := header.last
	for i := 0; i < header.ncells; i++ {
		cell.removeRow()
		cell = cell.down
	}
}

// Add this column and all the rows it uses.
func (header *Header) AddColumn() {
	// Add all the rows going up.
	cell := header.last.top
	for i := 0; i < header.ncells; i++ {
		cell.addRow()
		cell = cell.top
	}

	header.right.left = header
	header.left.right = header
}

// Remove the whole row except this cell, going right.
func (cell *Cell) removeRow() {
	for c := cell.right; c != cell; c = c.right {
		c.header.ncells--
		c.top.down = c.down
		c.down.top = c.top
		c.header.last = c.top // This is a bit dodgy :/
	}
}

// Add the whole row except this cell going left.
func (cell *Cell) addRow() {
	for c := cell.left; c != cell; c = c.left {
		c.header.ncells++
		c.top.down = c
		c.down.top = c
	}
}

func (m *SparseMatrix) Solvable() bool {
	if m.iteration > 100 {
		return false
	}
	m.iteration++
	if m.head == nil {
		return true
	}
	header := getSmallestColumn(m.head)
	if header.ncells == 0 {
		return false
	}

	selectedCell := header.last
	for i := 0; i < header.ncells; i++ {
		selectedCell.RemoveAllAffectedColumns()

		if m.Solvable() {
			return true
		}
		selectedCell.AddAllAffectedColumns()
		selectedCell = selectedCell.down
	}
	return false
}

func (m *SparseMatrix) FixValue(row, column, value int) {
	index := getValueConstraintIndex(row, column, m.size)
	code := AffectationToCode(row, column, value, m.size)

	header := m.headers[index]
	cell := header.last
	for i := 0; i < header.ncells; i++ {
		if cell.code == code {
			cell.RemoveAllAffectedColumns()
		}
		cell = cell.down
	}
	fmt.Println("Couldn't find the right cell for affectation (%i, %i) = %i", row, column, value)
}
