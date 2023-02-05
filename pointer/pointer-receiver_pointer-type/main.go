package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape1 interface {
	area() float64
}

type shape2 interface {
	perimeter() float64
}

func main() {
	c1 := circle{5}
	fmt.Printf("c1.area(): %v\n", c1.area())
	fmt.Printf("(&c1).area(): %v\n", (&c1).area())
	fmt.Printf("c1.perimeter(): %v\n", c1.perimeter())
	fmt.Printf("(&c1).perimeter(): %v\n", (&c1).perimeter())
	//info1(c1) //not work
	info1(&c1)
	info2(c1)
	info2(&c1)

}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return math.Pi * c.radius * 2
}

func info1(s shape1) {
	fmt.Printf("s area: %v\n", s.area())
}

func info2(s shape2) {
	fmt.Printf("s perimeter: %v\n", s.perimeter())
}
