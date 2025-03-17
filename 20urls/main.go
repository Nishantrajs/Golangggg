package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const myurl string = "https://github.com/Nishantrajs"

func main() {

	fmt.Println("Welcome to Handling URLs in Golang!!!")

	// parsing

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.Request.URL.Scheme)
	fmt.Println(response.Request.URL.Host)
	fmt.Println(response.Request.URL.Path)
	fmt.Println(response.Request.URL.Port())
	fmt.Println(response.Request.URL.RawQuery)

	qparams := response.Request.URL.Query()

	fmt.Printf("Tye type of query params are %T\n", qparams)

	for _, val := range qparams {

		fmt.Println("\nParam is: ", val)
	}

	partsOfURL := &url.URL{

		Scheme:  "https",
		Host:    "https://github.com",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL)
}
