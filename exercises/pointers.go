package main

import "fmt"

func Swap(a, b *int) {
	*a, *b = *b, *a
}
func SwapBad(a, b int) {
	a, b = b, a
}
func main() {
	x, y := 1, 2
	Swap(&x, &y)
	fmt.Printf("%v %v", x, y)
}
