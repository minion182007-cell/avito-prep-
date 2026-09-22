package main

import "fmt"

func Apply(nums []int, f func(int) int) []int {
	nums1 := make([]int, len(nums))
	for j, i := range nums {
		nums1[j] = f(i)
	}
	return nums1
}
func main() {
	double := func(x int) int { return x * 2 }
	fmt.Print(Apply([]int{1, 2, 3}, double))
	c := Counter()
	fmt.Println(c())

	fmt.Println(c())

	fmt.Println(c())

	fmt.Println(c())
	d := Counter()
	fmt.Println(c())

	fmt.Println(d())

}
func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
