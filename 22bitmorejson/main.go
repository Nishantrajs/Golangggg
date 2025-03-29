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
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(` 
	{
                "coursename": "ReactJs BootCamp",
                "Price": 299,
                "Platform": "LearnCodeOnline.in",
                "tags": ["web-dev","js"]
        }
		`)

	var lcoCourses course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("JSON was not valid!!!")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is : %T\n", k, v, v)
	}
}
