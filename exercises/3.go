package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < int(math.Pow(float64(n), 0.5)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(isPrime(a))

}
