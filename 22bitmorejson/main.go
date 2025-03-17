package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json: "website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video in Golang !!!")
	EncodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJs BootCamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"Mern BootCamp", 299, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular BootCamp", 299, "LearnCodeOnline.in", "pqr123", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
