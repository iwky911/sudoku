package main

import(
	"github.com/iwky911/sudoku/tools"
	"fmt"
)

func main(){
	originalMatrix, _ = tools.ParseInput()
	createMatrix()
	
	if solvable(getSmallerColumn(m.head)) {
		fmt.Println("sudoku solvable !")
		for _,t := range originalMatrix {
			fmt.Println(t)
		}
	}else{
		fmt.Println("Sorry, this soduko is impossible")
	}
}
