package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Measurable interface {
	Perimeter() float64
}

type Geometry interface {
	Shape
	Measurable
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return r.width + r.width + r.height + r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	rectangle := Rectangle{width: 5, height: 4}
	circle := Circle{radius: 2}

	fmt.Println("Rectangle area:", calculateArea(rectangle))
	fmt.Println()
	fmt.Println("Circle area:", calculateArea(circle))
	fmt.Println()
	describeShape(rectangle)
	fmt.Println()
	describeShape(circle)
}

func describeShape(g Geometry) {
	fmt.Println("Area:", g.Area())
	fmt.Println("Perimeter:", g.Perimeter())
}

func calculateArea(s Shape) float64 {
	return s.Area()
}
