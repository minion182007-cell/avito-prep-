package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	a := strings.Fields(s)
	seen := make(map[string]int)
	for _, i := range a {
		seen[i] += 1
	}
	return seen
}
func main() {
	var s string
	s = "go go gopher"
	fmt.Println(WordCount(s))
	m := map[string]int{"a": 1}
	fmt.Println(m["b"])

	v, ok := m["b"]
	fmt.Println(v, ok)

	v, ok = m["a"]
	fmt.Println(v, ok)
	ages := map[string]int{
		"ann":   23,
		"boris": 31,
		"clara": 27,
		"dima":  19,
		"eva":   35,
	}
	for k, v := range ages {
		fmt.Println(k, v)
	}
}
