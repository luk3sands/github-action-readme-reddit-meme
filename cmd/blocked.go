package cmd

import (
	"os"
	"strings"
)

func LoadBlockedList() []string {

	// check if we have a block file
	_, err := os.Stat(".github/reddit/block.txt")
	if os.IsNotExist(err) {
		return []string{}
	}

	block, err := os.ReadFile(".github/reddit/block.txt")
	if err != nil {
		panic(err.Error())
	}

	return strings.Split(string(block), "\n")
}
