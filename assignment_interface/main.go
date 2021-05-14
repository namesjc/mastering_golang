package main

import "fmt"

type shape interface {
	getArea() float64
}


type triangle struct {
	base float64
	height float64
}

type square struct {
	sideLength float64
}


func main() {
	t := triangle{
		base: 3,
		height: 4,
	}

	s := square{
		sideLength: 5,
	}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64{
	area := 0.5 * t.base * t.height
	return area
}

func (s square) getArea() float64{
	area := s.sideLength * s.sideLength
	return area
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}


