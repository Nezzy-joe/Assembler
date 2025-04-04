package main

import (
	"fmt"
	"os"
)

func writeFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString("Hello Life from Nezzy")
	if err != nil {
		panic(err)
	}
}

func readFile() {
	data, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func main() {
	writeFile()
	readFile()
}
