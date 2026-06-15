package main

import (
	"fmt"
	"math"
)

// define a shape interface
type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{0, 0, 5}
	rectangle := Rectangle{10, 5}

	fmt.Printf("Circle Area: %.3f\n", getArea(circle))
	fmt.Printf("Circle Perimeter: %.3f\n", circle.perimeter())
	fmt.Printf("Rectangle Area: %.3f\n", getArea(rectangle))
	fmt.Printf("Rectangle Perimeter: %.3f\n", rectangle.perimeter())
}
