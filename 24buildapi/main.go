package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Model for course - file

type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to APIs in Golang !!")

}

// Controllers -- file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1> Welcome to API by learnCodeOnline !</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get all courses")
	// sets the response writer as content type and tells the client that can be web browser or postman that the file or response is sent in the json/format..
	w.Header().Set("Content-Type", "application/json")
	//Now the json format is now encoded to give you file in the form of json which we have created in fake db. This process is called as seeding.
	json.NewEncoder(w).Encode(courses)
}
