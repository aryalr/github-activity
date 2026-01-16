package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

// Blueprint
type Repo struct {
	Name string `json:"name"`
}

type Event struct {
	Type string `json:"type"`
	Repo Repo   `json:"repo"`
}

// Link Constructor
func link(username string) string {
	endpoint := "https://api.github.com/users/"
	return endpoint + username + "/events"
}

// Get data from user. (slice)
func githubLink() string {
	if len(os.Args) < 2 {
		return "Username tidak boleh kosong. contoh: github-activity <your_username>"
	}
	var username = os.Args[1]
	var github = link(username)
	return github
}

// Fetch data from githubLink()
func fetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return (bodyBytes), nil
}

func main() {
	var events []Event
	var link = githubLink()
	var data, err = fetchData(link)
	if err != nil {
		println(err)
		return
	}
	json.Unmarshal(data, &events)
	for _, event := range events {
		println("Event: " + event.Type + " di Repo: " + event.Repo.Name)
	}

}
