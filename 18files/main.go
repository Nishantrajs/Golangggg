package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to Files in Golang !!!")

	content := "This needs to go in a file - Golang.in"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {

		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {

		panic(err)
	}

	fmt.Println(length)
	readFile("./mylcogofile.txt")
	defer file.Close()
}

func readFile(fileName string) {

	databyte, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is \n", string(databyte))
}
