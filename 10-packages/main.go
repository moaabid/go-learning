package main

import (
	"fmt"

	"github.com/moaabid/learning/randomstuff"
	"github.com/moaabid/learning/shapes"
)

func main() {

	circle := shapes.Circle{Radius: 5}

	rectangle := shapes.Rectangle{Width: 10, Height: 5}

	fmt.Println(circle.Area())

	fmt.Println(rectangle.Area())

	fmt.Println(rectangle.IsSquare())

	randomstuff.PrintRandomStuff()
}
