package main

import "fmt"

func main() {

	var myFavoriteFood string

	myFavoriteFood = "pizza"

	fmt.Println(myFavoriteFood)

	changeMyFavoriteFood(&myFavoriteFood, "burger")

	fmt.Println(myFavoriteFood)
}

func changeMyFavoriteFood(fruitePointer *string, newFruit string) {
	fmt.Println("Address of my favorite food:", fruitePointer, "value of my favorite food:", *fruitePointer)

	*fruitePointer = newFruit
}
