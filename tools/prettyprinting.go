package tools

import (
	"fmt"
	"strconv"
	"strings"
)

func PrintSolutionFromCode(sol []int, size int) {
	matrix := make([][]string, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]string, size)
	}

	for _, code := range sol {
		r := code / (size * size)
		c := (code / size) % size
		v := code%size + 1
		matrix[r][c] = strconv.Itoa(v)
	}

	for _, row := range matrix {
		fmt.Println(strings.Join(row, ","))
	}
}
