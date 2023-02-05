package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func main() {

	c := circle{5}
	fmt.Println("area", c.area())
	fmt.Println("area", (&c).area())
	//info(c) //not work, pointer receiver only accept pointer value
	info(&c)

}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Printf("s area: %v\n", s.area())
}
