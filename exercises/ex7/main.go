package main

import "fmt"

func Sslise(a, b []int) []int {
	var NewSlise []int
	for _, x := range a {
		for _, y := range b {
			if x == y {
				NewSlise = append(NewSlise, x)
				break
			}

		}
	}
	return NewSlise
}
func Sslise1(a, b []int) []int {
	NewSlise := make([]int, 0, len(a)+len(b))
	seen := make(map[int]bool)
	for _, x := range a {
		if !seen[x] {
			seen[x] = true
			NewSlise = append(NewSlise, x)

		}

	}
	for _, y := range b {
		if !seen[y] {
			seen[y] = true
			NewSlise = append(NewSlise, y)
		}
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
	var n1 int
	fmt.Scan(&n1)
	s2 := make([]int, n1)
	for i := 0; i < n1; i++ {
		fmt.Scan(&s2[i])
	}
	fmt.Println(Sslise(s, s2))
	fmt.Println(Sslise1(s, s2))

}
