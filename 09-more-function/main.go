package main

import "fmt"

func main() {

	fmt.Println(calculateAverage(1.9, 2.7, 3.0, 4.5, 5.7))
	boringFunction("Hello", "World")

	func() {
		fmt.Println("I'm a function")
	}()

	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))

	hundred := 100

	addHundreds := func() int {
		hundred += 100
		return hundred
	}

	fmt.Println(addHundreds())

}

func boringFunction(someBoringStuff ...interface{}) {
	fmt.Println(someBoringStuff...)
}

func calculateAverage(numbers ...float64) float64 {

	var total float64
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}
