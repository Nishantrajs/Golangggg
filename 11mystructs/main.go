package main

import "fmt"

func main() {

	fmt.Println("Welcome to Structs in Golang !!")

	// no inheritance in Golangg
	// no super
	// no parent

	nishant := User{"Nishant", "nishantraj839@gmail.com", true, 26}

	fmt.Println(nishant)

	fmt.Printf("Nishant Details are : %+v\n", nishant)
	fmt.Printf("Name is %v and email is %v.", nishant.Name, nishant.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
