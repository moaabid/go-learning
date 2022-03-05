package main

import (
	"fmt"
	"sort"
)

func main() {

	//Array

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

	//Slices
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

	fmt.Println("Average score is: ", averagerScore)

}

func averageScore(scores []int) float64 { 
	var total int
	for _, score := range scores {
		total += score
	}
	return float64(total / len(scores))
}
