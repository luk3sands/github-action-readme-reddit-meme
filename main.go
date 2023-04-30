package main

import "os"

func main() {
	println("Updating code")

	file, _ := os.ReadFile("main.go")
	println(string(file))
}
