package main

import (
	"flag"
	"fmt"
	"github.com/iwky911/sudoku/dancinglinks"
	"github.com/iwky911/sudoku/tools"
	"math/rand"
	"time"
)

var sizeFlag = flag.Int("n", 9, "Size of the sudoku to generate")
var uncoveredFlag = flag.Float64("-uncovered", 0.5,
	"Percentage of the cells that are uncovered")

func main() {
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())
	size := *sizeFlag
	m := dancinglinks.NewSparseMatrix(size)

	possibleCases := rand.Perm(size * size)
	solution := make([]int, 0)
	for _, c := range possibleCases[:size/2] {
		value := rand.Intn(size)
		code, err := m.FixValue(c/size, c%size, value)
		if err != nil {
			fmt.Println("There was a conflict. Try again :)")
			return
		}
		solution = append(solution, code)
	}

	possible, s := m.GetSolution()
	if !possible {
		fmt.Println("The generated sudoku didn't have a solution. Try again :)")
	}

	// Now that we have the solution, shuffle it.
	solution = append(solution, s...)
	perm := rand.Perm(len(solution))
	partialSol := make([]int, len(solution))
	for i, v := range perm {
		partialSol[i] = solution[v]
	}

	// Select only a part of the solution to be visible.
	solutionSize := int(float64(len(solution)) * (*uncoveredFlag))

	tools.PrintSolutionFromCode(partialSol[:solutionSize], size)
}
