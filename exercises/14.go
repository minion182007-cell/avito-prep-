package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Perimetr() float64
	Area() float64
}
type Circle struct {
	rad float64
}
type Tri struct {
	A, B, C float64
}
type Squar struct {
	Weight, width float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.rad * c.rad
}
func (t Tri) Area() float64 {
	p := (t.A + t.B + t.C) / 2
	return math.Sqrt(p * (p - t.B) * (p - t.A) * (p - t.C))
}
func (s Squar) Area() float64 {
	return s.Weight * s.width
}
func (c Circle) Perimetr() float64 {
	return 2 * math.Pi * c.rad
}
func (t Tri) Perimetr() float64 {
	return t.A + t.B + t.C
}
func (s Squar) Perimetr() float64 {
	return (s.Weight + s.width) * 2
}
func describe(s Shape) {
	fmt.Printf("(%v %v)\n", s.Area(), s.Perimetr())
}

type SquarO struct {
	Op string
}

func (s SquarO) String() string {
	return fmt.Sprintf("%v ", s.Op)
}
func (s Squar) String() string {
	return fmt.Sprintf("%.2f ", s.Weight)
}

func main() {
	shapes := []Shape{
		Circle{rad: 5},
		Tri{A: 3, B: 4, C: 5},
		Squar{Weight: 3, width: 4},
	}
	fmt.Printf("%v %T", shapes[2], shapes[2])
	for _, s := range shapes {
		describe(s)
	}
}
