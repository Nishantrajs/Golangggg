package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to Slices in Golang !!")

	var fruitList = []string{"Apple", "Tomato", "Peach"}

	fmt.Printf("Type of data of fruitList is: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 465
	highScores[3] = 867

	fmt.Println(highScores)

	var arr = []int{}

	arr = append(arr, 5, 3, 1, 4)
	fmt.Println("The array is: ", arr)
	fmt.Println("The Size of the array is: ", len(arr))

	//sort the array using indexes

	sort.Ints(arr)
	fmt.Println(arr)

	// how to remove a value from slieces based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}

	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
