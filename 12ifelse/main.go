package main

import "fmt"

func main() {

	fmt.Println("If else in Golang !!!")

	loginCount := 23

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 && loginCount < 25 {
		result = "User Exists!!"
	} else {
		result = "Something else"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even!!!")
	} else {
		fmt.Println("Number is odd!!!")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println(("Num is not less than 10"))
	}
}
