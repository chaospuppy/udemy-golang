package main

import "fmt"

type geo interface {
	getArea() float64
}

type square struct {
	length float64
}

type triangle struct {
	height float64
	length float64
}

func main() {
	s := square{length: 4}
	t := triangle{height: 2, length: 4}

	printArea(s)
	printArea(t)
}

func printArea(g geo) {
	area := g.getArea()
	fmt.Printf("Area is: %f\n", area)
}

func (s square) getArea() float64 {
	area := s.length * 2
	return area
}

func (t triangle) getArea() float64 {
	area := t.length * t.height * 0.5
	return area
}
