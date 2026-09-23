package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Rect struct {
	W, H float64
}
type Circle struct {
	R float64
}

func (r *Rect) Area() float64 {
	return r.W * r.H
}
func (r Circle) Area() float64 {
	return math.Pi * r.R * r.R
}
func TotalArea(shapes []Shape) float64 {
	sum := 0.0
	for _, v := range shapes {
		sum += v.Area()
	}
	return sum
}
func (r Rect) String() string {
	return fmt.Sprintf("Rect(%gx%g)", r.W, r.H)
}
func main() {
	shapes := []Shape{
		&Rect{3, 4},
		Circle{1},
		&Rect{2, 2},
	}
	fmt.Println(TotalArea(shapes))
	fmt.Println(Rect{3, 4})
}
