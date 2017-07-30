package tools

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Affectation struct {
	Row, Column, Value int
}

// Expecting a format:
// size of the array, then list of numbers. -1 means no affectation.
func ParseInput() ([][]int, int) {
	var n int
	fmt.Scan(&n)
	tab := make([][]int, n)
	cells := make([]int, n*n)
	for index := range tab {
		tab[index], cells = cells[:n], cells[n:]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&tab[i][j])
		}
	}
	return tab, n
}

//func ParseCSVInputAsTab(reader *io.Reader) ([][]int, int) {

//}

func ParseCSVInput(input io.Reader) ([]Affectation, int, error) {
	reader := csv.NewReader(input)
	values, err := reader.ReadAll()

	if err != nil {
		return nil, 0, err
	}
	size := len(values)
	for _, row := range values {
		if len(row) != size {
			return nil, 0, errors.New("A row has the wrong number of elements")
		}
	}
	affectations := make([]Affectation, 0)
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			v, err := strconv.Atoi(values[r][c])
			if err == nil {
				if v <= 0 || v > size {
					return nil, 0, errors.New("Incorrect value found:")
				}
				affectations = append(affectations, Affectation{r, c, v - 1})
			}
		}
	}

	return affectations, size, nil
}
