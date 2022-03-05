package main

import "fmt"

type Person struct {
	firstName, lastName, country string
	age                          int
}

type Book struct {
	title, author string
	pubYear       int
}

func main() {

	personOne := Person{
		firstName: "James",
		lastName:  "Bond",
		country:   "United States",
		age:       32,
	}

	bookOne := Book{
		title:   "Casino Royale",
		author:  "Ian Fleming",
		pubYear: 1953,
	}

	fmt.Println("person one is: ", personOne)

	fmt.Println("person one full name is: ", personOne.getFullName())

	fmt.Println("book one title is: ", bookOne.title)

	bookOne.changeTitle("Casino Royale 2")

	fmt.Println("book one title is: ", bookOne.title)

}

// changeTitle is a method on Book
func (book *Book) changeTitle(newTitle string) {
	book.title = newTitle
}

// getFullName is a method on Person
func (person *Person) getFullName() string {
	return person.firstName + " " + person.lastName
}
