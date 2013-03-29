package main

import(
	"github.com/iwky911/sudoku/tools"
	"fmt"
)

func main(){
	originalMatrix, _ = tools.ParseInput()
	//fmt.Println(originalMatrix)
	createMatrix()
	
	fmt.Println(solvable(smallestNbCells(m.head)))
	for _,t := range originalMatrix {
			fmt.Println(t)
		}
}
