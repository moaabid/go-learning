package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {

	//Array
	/*
		-> Arrays are a collection of values of the same type.
		-> Arrays are zero-indexed.
		-> Arrays are fixed in size.
	*/

	//Array of 4 string
	var cars = [4]string{
		"Volvo",
		"BMW",
		"Ford",
		"Mazda",
	}

	fmt.Println("First car: ", cars[0])
	fmt.Println("length of array:", len(cars))
	cars[1] = "Toyota"
	fmt.Println("Second car: ", cars[1])

	//Array with 4 elements of type string
	cars2 := [...]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars2)

	//Assigning values to specific indexes of an array
	cars3 := [...]string{1: "Volvo", 5: "Mazda"}
	fmt.Println(cars3)

	//two dimensional array
	var matrix = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("2 dimensional array : ", matrix)

	//three dimensional array
	var matrix2 = [3][3][3]int{
		{{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9}},
		{{10, 11, 12},

			{13, 14, 15},
			{16, 17, 18}},
		{{19, 20, 21},

			{22, 23, 24},
			{25, 26, 27}},
	}
	fmt.Println("3 dimensional array : ", matrix2)

	//Slices
	/*
		-> Slices are a reference type.
		-> Slices are a dynamically-sized, flexible view into the elements of an array.
		-> Appending to a slice modifies the existing data, not the underlying array.
		-> A slice does not store any data, it just describes a section of an underlying array.
	*/

	//Creating a slice foods
	var foods []string

	foods = append(foods, "apple")
	foods = append(foods, "banana")
	foods = append(foods, "orange")
	foods = append(foods, "grapes")

	fmt.Println("First food: ", foods[0])
	fmt.Println("length of slice:", len(foods))
	fmt.Println("foods: ", foods)

	foods[2] = "pineapple"
	fmt.Println("Third food: ", foods[2])

	
	fmt.Println("foods: ", foods)

	//sort slice
	sort.Strings(foods)
	fmt.Println("Sorted foods: ", foods)

	// for i := 0; i < len(foods); i++ {
	// 	fmt.Println("Index", i, "value: ", foods[i])
	// }

	for index, foods := range foods {
		fmt.Println("Index", index, "value: ", foods)
	}

	var scores = []int{1, 2, 3, 4, 5}

	averagerScore := averageScore(scores)
	log.SetFlags(log.Lshortfile)
	log.Println(averagerScore)
	fmt.Println("Average score is: ", averagerScore)

}

func averageScore(scores []int) float64 {
	var total int
	for _, score := range scores {
		total += score
	}
	return float64(total / len(scores))
}
