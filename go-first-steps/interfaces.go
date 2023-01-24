package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}
type circ struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c circ) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}
func (c circ) perimeter() float64 {
	return 2 * c.radius * math.Pi
}

func measure(g geometry) {
	fmt.Println("Geometry -> ", g)
	fmt.Println("Area -> ", g.area())
	fmt.Println("Perimeter -> ", g.perimeter())
}

func main() {
	r := rectangle{width: 10, height: 100}
	c := circ{2}

	measure(r)
	measure(c)
}
