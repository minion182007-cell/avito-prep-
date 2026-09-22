package main

import "fmt"

func printTable(n int) [][]int {
	tab := make([][]int, n)
	for i := 0; i < n; i++ {
		tab[i] = make([]int, n)
		for j := 0; j < n; j++ {
			tab[i][j] = (i + 1) * (j + 1)
		}
	}
	return tab
}
func main() {
	var a int
	fmt.Scan(&a)
	tab := printTable(a)
	for _, row := range tab {
		for _, col := range row {
			fmt.Printf("%v\t", col)
		}
		fmt.Println()

	}

}
