package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func CommitReadme() {

	cmdAdd := exec.Command("git", "add", "README.md")
	_, err := cmdAdd.Output()
	if err != nil {
		fmt.Println("Error adding to git: ", err.Error())
		return
	}

	cmdCommit := exec.Command("git", "commit", "-m", "Change README.md")
	cmdCommit.Stdout = os.Stdout
	cmdCommit.Run()

	cmdPush := exec.Command("git", "push")
	cmdPush.Stdout = os.Stdout
	cmdPush.Run()
}
