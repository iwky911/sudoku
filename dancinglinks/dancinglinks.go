package main

import(
	"math"
	//"fmt"
	//"os"
)
/*
Contains the struct
* first n elt: validate row i contains v
* second n elts: validates column i contains v
* final n elts: validates block i contains v
*/

func (c *Cell) isCorrect() bool{
	return c.right!=nil && c.left!=nil
}

func columnOk(c *Cell) bool{
	if c.down == nil {return true}
	if c.down==c.down.down{
		return c.down.isCorrect()
	}
	for n:=c.down.down; n!=c.down; n=n.down{
		if !n.isCorrect(){
			return false
		}
	}
	return true
}

func columnPrint(c *Cell){
	if c.down==nil || c.down.down==c.down {
		//fmt.Println("one element")
	}
	for n:=c.down.down; n!=c.down; n=n.down{
		//fmt.Println("element: ",n)
	}
}
var m SparseMatrix

var originalMatrix [][]int

type SparseMatrix struct{
	heads []*Cell
	head *Cell
	n, ns int
}

// func (m *SparseMatrix) size()

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
	//fmt.Println(len(m.heads))
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
	////fmt.Println("inserting cell on column ",i, len(m.heads))
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
	if i==0 {
		//fmt.Println(m.heads[0].down)
		//fmt.Println(m.heads[0].down.down)
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

// func (c *Cell) removeLine() {

// 	// remove all links to this line except for this column
// 	for next:= c.right; next!= c; next=next.right {
// 		next.removeCell()
// 		//fmt.Println("cell removed: ",next.value)
// 	}
//}

func (c *Cell) restoreCell(){
	if c.top == c {
		// was alone: must relink head
		c.head.down = c
	}else{
		c.top.down = c
		c.down.top = c
	}
}

// func (c *Cell) restoreLine(){
// 	for next := c.right; next!=c; next = next.right {
// 		next.restoreCell()
// 	}
// }

func (c *Cell) removeColumn(){
	////fmt.Println(c)

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
	//fmt.Println(cells[0])
	//fmt.Println(cells[2])
	//fmt.Println(cells[0].value)
	addCell(cells[0], i*m.n + v)
	addCell(cells[1], m.n*m.n + j*m.n + v)
	addCell(cells[2], m.n*m.n*2 + m.n*((i/m.ns)*m.ns + j/m.ns) + v)
	addCell(cells[3], m.n*m.n*3 + i*m.n+j)
}

func nbCells(col *Cell) int {
			// if !columnOk(col){
			// 	//fmt.Println("error!")
			// }
	if col.down==nil {
		return 0
	}
	if col.down.down == col.down {
		return 1
	}
	i:=1
	for elt:= col.down.down; elt!=col.down; elt= elt.down {
		i++
		//if i>1000 {//fmt.Println(elt,col.down, elt!=col.down, i)}
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
		//fmt.Println("after nbcell", c.down.value,"\n", col.down.value)
		if nb<rec {
			rec=nb
			index=c
		}
		////fmt.Println("echo")
	}
	//fmt.Println("smallest : ",rec)
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
	// for cell:= sel.right; cell!=sel; cell=cell.right{
	// 	cell.removeCell()
	// }
}

func (sel *Cell) restoreLinked(){
	//fmt.Println("restoring")
	sel.restoreColumn()
	for col:= sel.right; col!=sel; col=col.right {
		col.restoreColumn()
		for row := col.down; row!=col; row=row.down{
			for cell:= row.right; row!=cell; cell = cell.right{
				cell.restoreCell()
			}
		}
	}
	// for cell:= sel.right; cell!=sel; cell=cell.right{
	// 	cell.restoreCell()
	// }
}

var depth = 0

func solvable(c *Cell) bool {
	// //fmt.Println("calling solvable ",depth)
	// for _,h := range m.heads {
	// 	if !columnOk(h){
	// 		//fmt.Println("error! for col ",h)
	// 	}
	// }
	// for i:=0; i<324; i++{
	// 		if m.heads[i].down==nil {
	// 			//fmt.Println(i)
	// 		}
	// 	}
	// return false
	switch{
	case c == nil:
		return true
	case c.down == nil:
		for _,v := range m.heads {
			if v==c {
				//fmt.Println("head empty at :",i)
			}
		}
		////fmt.Println("here")
		// TODO: delete that
		
		return false
	default:
		// //fmt.Println("before")
		// //fmt.Println("after")
		// //fmt.Println("col correct: ", columnOk(c))
		first:=true
		var next *Cell
		for elt :=c.down; (elt != c.down) || (elt == c.down && first) ; elt = elt.down {
			//fmt.Println("affecting !",elt.value)
			////fmt.Println("heads ok: ",headsOk(c))
			depth++
			// //fmt.Println(depth)
			
			// //fmt.Println("removing column :",c)
			// // //fmt.Println("solving for ",next)
			// for column in selectedrow {
			// 	for row in col {
			// 		remove row
			// 	}
			// 	remove col
			// }

			// for col:= elt.right; col:=elt; col=col.right {
			// 	for row := col.down; row!=col; row=row.down{
			// 		row.removeLine()
			// 	}
			// }
			// elt.removeColumn()
			// elt.removeLine()
			// for node := elt.right; node!= elt; node = node.right{
			// 	node.removeColumn()
			// 	//fmt.Println("heads touched: ",node.head)
			// }
			elt.removeLinked()
			//fmt.Println("after remove\n", c,"\n", c.right,"\n", c.right.right)
			//for 
			// if c.right==c {
			// 	next=nil
			// }else{
				//fmt.Println("here",elt.head.right.down.value)
				next = smallestNbCells(m.head)
			// }
			//fmt.Println("powerthrough ",depth)
			if solvable(next){
				aff := elt.value
				originalMatrix[aff[0]][aff[1]]=aff[2]
				return true
			}
			//fmt.Println("backup")
			depth--
			elt.restoreLinked()
			//fmt.Println("after restore")
			// elt.restoreLine()
			// elt.restoreColumn()
			first=false
		}
		return false
	}
	return false
}


