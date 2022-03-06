package main

import "fmt"

type User struct {
	firstName, lastName string
	Age                 int
}

func main() {

	graded := make(map[string]int)

	graded["Timmy"] = 75
	graded["Jess"] = 88
	graded["Sam"] = 99

	fmt.Println(graded["Jess"])

	userData := make(map[string]User)

	userData["Timmy"] = User{
		firstName: "Timmy",
		lastName:  "Turner",
		Age:       25,
	}

	userData["Jess"] = User{ 
		firstName: "Jess",
		lastName:  "Johnson",
		Age:       22,
	}

	userData["Sam"] = User{
		firstName: "Sam",
		lastName:  "Smith",
		Age:       30,
	}

	//CAUTION: Without astricks, cannot change the value of the key
	//userData["Sam"].firstName = "Sammy"

	fmt.Println(userData)

	userDataBase := make(map[string]*User)

	userOne := User{
		firstName: "Timmy",
		lastName:  "Turner",
		Age:       25,
	}

	userTwo := User{
		firstName: "Jess",
		lastName:  "Johnson",
		Age:       22,
	}
	userThree := User{
		firstName: "Sam",
		lastName:  "Smith",
		Age:       30,
	}

	userDataBase["Timmy"] = &userOne
	userDataBase["Jess"] = &userTwo
	userDataBase["Sam"] = &userThree
	fmt.Println(userDataBase["Sam"])

	//We can change the value of the key
	userDataBase["Sam"].firstName = "Sammy"

	fmt.Println(userDataBase["Sam"].firstName)
}
