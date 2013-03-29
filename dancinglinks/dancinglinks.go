package main

import(
	"math"
)
/*
Contains the struct
* first n elt: validate row i contains v
* second n elts: validates column i contains v
* final n elts: validates block i contains v
*/
var m SparseMatrix

var originalMatrix [][]int

type SparseMatrix struct{
	heads []*Cell
	head *Cell
	n, ns int
}


type Cell struct{
	top, down, left, right, head *Cell
	value *[3]int
}

func createMatrix() {
	n:= len(originalMatrix)
	size:=4*n*n
	m = SparseMatrix{}
	m.heads = make([]*Cell, size)
	m.n= n
	m.ns= int(math.Sqrt(float64(n)))
	//creating heads
	for i:=0; i<size; i++ {
		m.heads[i] = &Cell{}
	}
	//creating links between heads
	for i:=0; i<size; i++ {
		m.heads[i].left = m.heads[(i-1+size)%size]
		m.heads[i].right = m.heads[(i+1)%size]
	}
	m.head = m.heads[0]
	for i:=0; i<m.n; i++{
		for j:=0; j<m.n; j++{
			if originalMatrix[i][j]!= -1 {
				addAffectation(i,j,originalMatrix[i][j]-1)
			}else{
				for v:=0; v<m.n; v++ {
					addAffectation(i,j,v)
				}
			}
		}
	}
}

func addCell(c *Cell, i int) {
	c.head = m.heads[i]
	if m.heads[i].down == nil{
		m.heads[i].down = c
		c.top=c
		c.down=c
	}else{
		before := m.heads[i].down
		after := m.heads[i].down.down
		before.down = c
		c.top = before
		c.down = after
		after.top = c
	}
}

func (c *Cell) removeCell() {
	if c.top == c {
		//alone in this column
		c.head.down=nil
	}else{
		c.top.down = c.down
		c.down.top = c.top
		c.head.down = c.down
	}
}

func (c *Cell) restoreCell(){
	if c.top == c {
		// was alone: must relink head
		c.head.down = c
	}else{
		c.top.down = c
		c.down.top = c
	}
}

func (c *Cell) removeColumn(){
	if c.head.left == c.head {
		// only column in the matrix !
		m.head = nil
	}else{
		if m.head==c.head{
			m.head = c.head.right
		}
		// removing links in header
		c.head.right.left = c.head.left
		c.head.left.right = c.head.right
	}
}

func (c *Cell) restoreColumn(){
	// restoring header links
	c.head.right.left = c.head
	c.head.left.right = c.head
	if m.head == nil{
		m.head=c
	}
}

func addAffectation(i,j,v int) {
	cells := [4]*Cell{}
	for i:=0; i<4; i++{
		cells[i] = &Cell{}
	}
	enr := [3]int{i,j,v}
	for i:=0; i<4; i++{
		cells[i].left = cells[(i+3) % 4]
		cells[i].right = cells[(i+1) % 4]
		cells[i].value = &enr
	}
	addCell(cells[0], i*m.n + v)
	addCell(cells[1], m.n*m.n + j*m.n + v)
	addCell(cells[2], m.n*m.n*2 + m.n*((i/m.ns)*m.ns + j/m.ns) + v)
	addCell(cells[3], m.n*m.n*3 + i*m.n+j)
}

func nbCells(col *Cell) int {
	if col.down==nil {
		return 0
	}
	if col.down.down == col.down {
		return 1
	}
	i:=1
	for elt:= col.down.down; elt!=col.down; elt= elt.down {
		i++
	}
	return i
}

func smallestNbCells(col *Cell) *Cell {
	if col==nil{
		return nil
	}
	rec := nbCells(col)
	index := col
	for c:= col.right; c!=col; c=c.right {
		nb:= nbCells(c)
		if nb<rec {
			rec=nb
			index=c
		}
	}
	return index
}

func (sel *Cell) removeLinked(){
	sel.removeColumn()
	for col:= sel.right; col!=sel; col=col.right {
		col.removeColumn()
		for row := col.down; row!=col; row=row.down{
			for cell:= row.right; row!=cell; cell = cell.right{
				cell.removeCell()
			}
		}
	}
}

func (sel *Cell) restoreLinked(){
	sel.restoreColumn()
	for col:= sel.right; col!=sel; col=col.right {
		col.restoreColumn()
		for row := col.down; row!=col; row=row.down{
			for cell:= row.right; row!=cell; cell = cell.right{
				cell.restoreCell()
			}
		}
	}
}

func solvable(c *Cell) bool {
	switch{
	case c == nil:
		return true
	case c.down == nil:
		return false
	default:
		first:=true
		for elt :=c.down; (elt != c.down) || (elt == c.down && first) ; elt = elt.down {
			elt.removeLinked()
			if solvable(smallestNbCells(m.head)){
				aff := elt.value
				originalMatrix[aff[0]][aff[1]]=aff[2]
				return true
			}
			elt.restoreLinked()
			first=false
		}
		return false
	}
	return false
}