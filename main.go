package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/d3code/github-action-commit-workflow-changes/cmd"
)

func main() {
	println("Finding some lols...")

	// Change to the workspace directory
	workspace := os.Getenv("GITHUB_WORKSPACE")
	if workspace != "" {
		os.Chdir(workspace)
	}

	// Get the image url
	var imageUrl string = cmd.SelectImage()
	if imageUrl == "" {
		fmt.Println("No post found with image")
		return
	}

	// Get the readme
	readme, err := os.ReadFile("README.md")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Find line with the image
	lines := strings.Split(string(readme), "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "![Reddit]") {
			lines[i] = fmt.Sprintf("![%s](%s)", "Reddit", imageUrl)
			break
		}
	}

	// Update the readme
	readme = []byte(strings.Join(lines, "\n"))
	err = os.WriteFile("README.md", readme, 0644)
	if err != nil {
		panic(err.Error())
	}

	// Commit the changes
	if os.Getenv("GITHUB_ACTION") != "" {
		cmd.CommitReadme()
	}

	println("Done!")
}
