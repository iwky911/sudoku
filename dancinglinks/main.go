package main

import (
	"flag"
	"fmt"
	"github.com/iwky911/sudoku/tools"
	"io"
	"os"
)

var csvfile = flag.String("csv", "", "csv file to parse")

func main() {
	flag.Parse()
	var reader io.Reader
	var err error
	if csvfile != nil {
		reader, err = os.Open(*csvfile)
	}
	affectations, size, err := tools.ParseCSVInput(reader)
	if err != nil {
		fmt.Println("Error while parsing:", err)
	}
	fmt.Println("Parsed a matrix of size", size)

	m := createSparseMatrix(size)
	partialsol := make([]int, 0, len(affectations))
	for _, affectation := range affectations {
		fmt.Printf("Affecting, (%v, %v) = %v\n", affectation.Row, affectation.Column, affectation.Value)
		code, err := m.FixValue(affectation.Row, affectation.Column, affectation.Value)
		if err != nil {
			fmt.Printf("Failed to do the affectation")
			return
		}
		partialsol = append(partialsol, code)
	}
	solvable, sol := m.Solvable()
	sol = append(sol, partialsol...)
	if solvable {
		fmt.Println("sudoku is solvable!!")
		PrintSolutionFromCode(sol, size)
	} else {
		fmt.Println("sudoku is not solvable :(")
	}
}
