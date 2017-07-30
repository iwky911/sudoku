package tools

import (
	"fmt"
	"sort"
)

func PrintSolutionFromCode(sol []int, size int) {
	sort.Ints(sol)

	for i, v := range sol {
		if i%size == 0 {
			fmt.Println()
		}
		fmt.Printf("%v,", (v%size)+1)
	}
	fmt.Println()
}
