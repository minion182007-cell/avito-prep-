package main

import "fmt"

func reverseN(s []int) []int {
	NewSlise := make([]int, len(s))
	for i, v := range s {
		NewSlise[len(s)-i-1] = v
	}
	return NewSlise
}
func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	fmt.Println(reverseN(s))

}
