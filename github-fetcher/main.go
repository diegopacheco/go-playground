package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Repo struct {
	Name string `json:"full_name"`
}

func extractRepos() []Repo {
	url := "https://api.github.com/orgs/Netflix/repos"
	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	repos := make([]Repo, 0)
	jsonErr := json.Unmarshal(body, &repos)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return repos
}

func main() {
	repos := extractRepos()
	for _, r := range repos {
		fmt.Println(r.Name)
	}
}
