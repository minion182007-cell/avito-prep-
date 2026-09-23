package main

import "fmt"

type Rect struct {
	W, H float64
}
type Celsius float64

func (r Rect) Area() float64 {
	return r.W * r.H

}
func (r *Rect) Scale(k float64) {
	r.W = r.W * k
	r.H = r.H * k

}
func (r Celsius) Fahrenheit() float64 {
	return float64(r)*9/5 + 32

}

func main() {
	r := Rect{3, 4}
	fmt.Println(r.Area())
	r.Scale(2)
	fmt.Println(r)
	fmt.Println(r.Area())
	c := Celsius(25)
	fmt.Println(c.Fahrenheit())
}
