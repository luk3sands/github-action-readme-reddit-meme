package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetReddit() Reddit {
	url := "https://www.reddit.com/r/ProgrammerHumor/.json"
	body := SendRequest(url)

	var response Reddit
	err := json.Unmarshal(body, &response)
	if err != nil {
		panic(err.Error())
	}

	return response
}

func SendRequest(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	return body
}
