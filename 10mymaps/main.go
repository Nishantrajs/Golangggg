package main

import (
	"fmt"
)

// func main() {

// 	fmt.Println("Welcome to Maps in Golangg!!")

// 	languages := make(map[string]string)

// 	languages["JS"] = "JavaScript"
// 	languages["RB"] = "Ruby"
// 	languages["PY"] = "Python"

// 	fmt.Println("List of all languages: ", languages)
// 	fmt.Println("JS Short for: ", languages["JS"])

// 	delete(languages, "RB")
// 	fmt.Println("List of all languages: ", languages)

// 	// loops are interesting in golang

// 	for key, value := range languages {
// 		fmt.Printf("For key %v, value is %v\n", key, value)
// 	}
// }

func main() {

	var arr = []int{1, 2, 1, 2, 3}

	var sum int = 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println(sum)

	var mp = make(map[int]int)
	result := solve(arr, mp)

	fmt.Println(result)

}

func solve(arr []int, mp map[int]int) int {

	var maxi int = -1

	for i := 0; i < len(arr); i++ {
		mp[arr[i]]++
	}

	fmt.Println(mp)

	for _, value := range mp {
		maxi = max(maxi, value)
	}

	return maxi
}
