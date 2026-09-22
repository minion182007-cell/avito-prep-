package main

import "fmt"

func Fibonahiter(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
func Fibonahirec(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * Fibonahirec(n-1)
	}
}

func Factorrec(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	} else {
		return Factorrec(n-1) + Factorrec(n-2)
	}
}
func main() {
	var a int
	fmt.Scanln(&a)
	fmt.Println(Fibonahiter(a))
	fmt.Println(Fibonahirec(a))
	fmt.Println(Factorrec(a))
}
