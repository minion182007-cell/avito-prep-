package main

import "fmt"

type Point struct {
	X, Y int
}

func Move(p *Point, dx, dy int) {
	p.X += dx
	p.Y += dy
}
func main() {
	p1 := Point{1, 2}
	p2 := p1
	p2.X = 100
	fmt.Println(p1.X)

	p3 := &p1
	p3.X = 100
	fmt.Println(p1.X)
}
