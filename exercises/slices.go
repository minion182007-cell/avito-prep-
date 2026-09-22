package main

import "fmt"

func Reverce(s []int) []int {
	t := make([]int, len(s))
	for j, i := range s {
		t[len(s)-j-1] = i
	}
	return t
}
func main() {
	var s []int
	for i := 0; i < 20; i++ {
		s = append(s, i)
		fmt.Printf("%v %v\n", len(s), cap(s))
	}
	fmt.Println(Reverce(s))
	a := []int{1, 2, 3, 4, 5}
	b := a[1:3]
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)
}
