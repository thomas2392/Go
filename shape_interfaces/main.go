package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	heigth float64
	base   float64
}

func main() {
	s := square{sideLength: 2}
	t := triangle{heigth: 2.5, base: 2}

	printArea(s)
	printArea(t)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.heigth
}

func printArea(s shape) {
	fmt.Println("Area is:", s.getArea())
}
