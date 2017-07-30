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
	for _, affectation := range affectations {
		fmt.Printf("Affecting, (%v, %v) = %v\n", affectation.Row, affectation.Column, affectation.Value)
		m.FixValue(affectation.Row, affectation.Column, affectation.Value)
	}
	if m.Solvable() {
		fmt.Println("sudoku is solvable!!")
	} else {
		fmt.Println("sudoku is not solvable :(")
	}
}
