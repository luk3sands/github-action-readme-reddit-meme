package main

import (
	"fmt"
	"os"
)

func main() {
	println("Updating code")

	file, err := os.ReadFile("main.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}
