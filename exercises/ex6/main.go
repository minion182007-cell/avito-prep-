package main

import "fmt"

func Dubl(s []int) []int {
	n := 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			s[n] = s[i]
			n++
		}
	}
	return s[:n]
}
func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	fmt.Println(Dubl(s))

}
