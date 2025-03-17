package main

import "fmt"

const LoginToken string = "newToken12343" // L is capital so - public variable

func main() {

	var username string = "Nishant Raj"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Varible is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println((smallVal))
	fmt.Printf("Variable is of type : %T \n", smallVal)

	var smallFloat float64 = 255.045464654
	fmt.Println((smallFloat))
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	var anotherVariable int = 5
	fmt.Println((anotherVariable))
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	// implicit type

	var website = "learncodeonline"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println((LoginToken))
	fmt.Printf("Variable is of type : %T \n", LoginToken)
}
