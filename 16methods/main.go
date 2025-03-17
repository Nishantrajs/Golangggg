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
	fmt.Printf("Name is %v and email is %v\n", nishant.Name, nishant.Email)

	nishant.GetStatus()
	nishant.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {

	fmt.Println("Is User active: ", u.Status)
}

func (u User) NewMail() {

	u.Email = "nishant1223209@gmail.com"
	fmt.Println("Email of this user is: ", u.Email)
}
