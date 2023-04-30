package cmd

func ArrayContainsString(array []string, findString string) bool {
	for _, arrayString := range array {
		if arrayString == findString && arrayString != "" {
			return true
		}
	}
	return false
}
