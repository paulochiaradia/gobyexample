package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
} //This area method has a receiver type of *rect.

func (r *rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{
		width:  10,
		height: 20,
	}
	fmt.Println("Area: ", r.area())
	fmt.Println("Perimeter: ", r.perimeter())

	rp := &r
	fmt.Println("Rp Area: ", rp.area())
	fmt.Println("Rp Perimeter: ", rp.perimeter())
}
