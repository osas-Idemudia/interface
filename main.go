package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base float64
}
type square struct {
	sideLength float64
}

func main() {
	tr := triangle{
		height: 4,
		base: 4,
	}
	sq := square{
		sideLength: 3,
	}

	printArea(tr)
	printArea(sq)
}

func (t triangle) getArea() float64 {
	return 0.5* t.base * t.height
}
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func printArea(s shape) {
	fmt.Println(s.getArea())
}
