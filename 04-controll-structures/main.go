package main

import "fmt"

func main() {

	var number int
	var favSport string

	number = 15

	//if statement
	if number%2 == 0 {
		fmt.Println("Even")
	} else if number%3 == 0 {
		fmt.Println("Divisible by 3")
	} else {
		fmt.Println("Odd")
	}

	//switch statement
	favSport = "soccer"

	switch favSport {
	case "cricket":
		fmt.Println("Cricket is my favorite sport")
	case "football":
		fmt.Println("Football is my favorite sport")
	case "soccer":
		fmt.Println("Soccer is my favorite sport")
	default:
		fmt.Println("I don't know that sport")
	}

	//for loop
	i := 1

	for i <= 10 {
		fmt.Println("i is: ", i)
		i++
	}

	for x := 1; x <= 20; x++ {
		fmt.Println("x is ", x)
	}
}
