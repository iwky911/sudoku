package tools

import(
	"fmt"
)

func ParseInput() ([][]int, int) {
	var n int
	fmt.Scan(&n)
	tab := make([][]int,n)
	cells := make([]int,n*n)
	for index:=range tab{
		tab[index], cells = cells[:n], cells[n:]
	}
	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			fmt.Scan(&tab[i][j])
		}
	}
	return tab, n
}