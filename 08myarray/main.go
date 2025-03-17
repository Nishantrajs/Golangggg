package main

import "fmt"

func main() {

	fmt.Println("Welcome to arrays in Golang !!")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Length is Fruit list is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veggy List is : ", vegList)
	fmt.Println("Length is Fruit list is: ", len(vegList))
}
