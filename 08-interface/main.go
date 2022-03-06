package main

import "fmt"

type Circle struct {
	radius float64
	name   string
}

type Rectangle struct {
	width, height float64
	name          string
}

type Shape interface {
	area() float64
	changeName(name string)
}

func (c *Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (r *Rectangle) area() float64 {
	return r.width * r.height
}

func main() {

	circle := Circle{radius: 5}

	rectangle := Rectangle{width: 100, height: 100}

	shape := []Shape{&circle, &rectangle}

	printInfo(shape)

}

func (c *Circle) changeName(name string) {
	c.name = name
}

func (r *Rectangle) changeName(name string) {
	r.name = name
}

func printInfo(shape []Shape) {
	for _, shape := range shape {
		fmt.Println(shape.area())

	}
}
