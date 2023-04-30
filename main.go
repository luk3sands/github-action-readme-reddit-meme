package main

import (
	"fmt"
	"os/exec"
)

func main() {

	command := exec.Command("git", "status")
	out, err := command.Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(out))
}
