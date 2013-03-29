package main

import(
	"fmt"
	"github.com/iwky911/sudoku/tools"
	"math"
)

var tableau [][]int

var n int
var ns int

func lineViolation(i,j int) bool {
	for k:=0; k<n; k++{
		if k!=j && tableau[i][j]==tableau[i][k] {
			return true
		}
	}
	return false
}

func rowViolation(i,j int) bool {
	for k:=0; k<n; k++{
		if k!=i && tableau[i][j]==tableau[k][j] {
			return true
		}
	}
	return false
}
func squareViolation(i,j int) bool{
	v:=tableau[i][j]
	is, js := (i/ns)*ns, (j/ns)*ns
	for a:=is;a<is+ns;a++{
		for b:=js;b<js+ns;b++{
			if v==tableau[a][b] && (a!=i || b!=j) {
				return true
			}
		}
	}
	return false
}

func nexts(i,j int) (a,b int) {
	if j<n-1 {
		return i,j+1
	}
	return i+1,0

}
/*
	return true if we can find a satifiable assignement for cell (i,j)
*/
func solvable(i,j int) bool {
	//if all cells are affected, we return
	if i==n {
		return true
	}

	if tableau[i][j]!=-1 {
		return solvable(nexts(i,j))
	}

	for guess:=1; guess<=n; guess++ {
		tableau[i][j] = guess
		if lineViolation(i,j) || rowViolation(i,j) || squareViolation(i,j) || !solvable(nexts(i,j)) {
			tableau[i][j]=-1
		}else{
			return true
		}
	}
	return false
}

func main(){
	tableau, n = tools.ParseInput()
	ns = int(math.Sqrt(float64(n)))
	if solvable(0,0) {
		fmt.Println("solvable! ")
		for _,t := range tableau {
			fmt.Println(t)
		}
	}else{
		fmt.Println("unsolvable")
	}
}