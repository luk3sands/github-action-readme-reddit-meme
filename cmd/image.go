package cmd

import (
	"fmt"
	"strings"
)

func SelectImage() string {

	reddit := GetReddit()
	blocked := LoadBlockedList()

	posts := reddit.Data.Children
	for _, post := range posts {

		fmt.Println(post.Data.Title)
		fmt.Println(post.Data.Url)

		if (strings.HasSuffix(post.Data.Url, ".jpg") ||
			strings.HasSuffix(post.Data.Url, ".gif") ||
			strings.HasSuffix(post.Data.Url, ".png")) &&
			!ArrayContainsString(blocked, post.Data.Url) {

			return post.Data.Url
		}
	}

	return ""
}
