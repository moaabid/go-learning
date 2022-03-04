package main

import (
	"fmt"
)

// Declare a global variable called global_variable
var global_variable = "global"

func main() {

	// Declare a variable called local_variable and assign the value "local" to it.
	var awesomeString string
	var reallyCoolInt int
	var nicebool = false
	var unsigned_int uint = 2
	var signed_int int = -2

	awesomeString = "Go is awesome!"
	reallyCoolInt = 42
	coolFloat := 3.14

	//list of variables can be declared in one line
	// var (
	// 	awesomeString2 string
	// 	reallyCoolInt2 int
	// 	nicebool2 = false
	// 	unsigned_int2 uint
	// 	signed_int2 int
	// )

	//declare multiple variables at once
	// var oneString, oneInt, oneBool, oneUint, oneInt2, oneBool2, oneUint2  = "one", 1, true, 2, -2, false, 3

	// Print the value of local_variable.
	fmt.Println("String : ", awesomeString)
	fmt.Println("Int : ", reallyCoolInt)
	fmt.Println("Float : ", coolFloat)
	fmt.Println("Boolean : ", nicebool)
	fmt.Println("Unsigned Int :", unsigned_int)
	fmt.Println("Signed Int : ", signed_int)

	// Print the value of global_variable.
	fmt.Println("Global Variable: ", global_variable)

	// Print the value of the function getFirstName().
	fmt.Println("firstname is : ", getFirstName())

	Mohammed, Aabid := getBunchOfNames()
	fmt.Println("name one : ", Mohammed)
	fmt.Println("name two : ", Aabid)

	//ignore first return value
	_, Aabid2 := getBunchOfNames()
	fmt.Println(Aabid2)

	//get all return values
	name, age, available := getDetail()
	fmt.Println("name : ", name)
	fmt.Println("age : ", age)
	fmt.Println("available : ", available)

	//incrementByOne(number)
	fmt.Println("incrementByOne : ", incrementByOne(1))

	//addTwoFloats(a, b)
	fmt.Println("addTwoFloats : ", addTwoFloats(1.1, 2.2))

	//getFunction() return a function and print Hello World
	getFunction()()

}

func getFirstName() string {
	return "Mohammed Aabid"
}

func getBunchOfNames() (string, string) {
	return "Mohammed", "Aabid"
}

func getDetail() (string, int, bool) {
	return "Mohammed Aabid", 23, true
}

func incrementByOne(number int) int {
	return number + 1
}

func addTwoFloats(a, b float64) float64 {
	return a + b
}

func getFunction() func() {
	return func() {
		fmt.Println("Hello World")
	}
}
