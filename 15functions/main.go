package main

import "fmt"

func main() {

	fmt.Println("Welcome to Functions in Golang !!!")
	greeter()

	var result int = solve(2, 3)

	fmt.Println("The Sum of Two Numbers is: ", result)

	var totalSum = proAdder(1, 2, 3)

	fmt.Println("The Sum of All values is: ", totalSum)
}

func solve(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int {
	var sum = 0

	for total := 0; total < len(values); total++ {
		sum += values[total]
	}

	return sum
}

func greeter() {

	fmt.Println("Namastey from Golang!!!")
}
