package main

import "fmt"

type rect struct {
	height, width int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) perimeter() int {
	return 2 * (r.height + r.width)
}

func main() {
	r := rect{width: 5, height: 10}
	fmt.Println("area: ", r.area())
	fmt.Println("perimeter: ", r.perimeter())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perimeter: ", rp.perimeter())
}
