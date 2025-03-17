package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"

func main() {

	fmt.Println("Welcome to Web Requests in Golang!!")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of Type %T\n", response)

	defer response.Body.Close() // Caller's responsibility is to close the connection

	dataByte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("The data is: ", string(dataByte))
}
