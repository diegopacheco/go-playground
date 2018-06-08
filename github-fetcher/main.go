package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Repo struct {
	Name string `json:"full_name"`
}

func extractRepos(page int) ([]Repo, error) {
	url := "https://api.github.com/orgs/Netflix/repos?page=" + strconv.Itoa(page)
	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error 1")
		return nil, err
	}
	res, getErr := client.Do(req)
	if getErr != nil {
		fmt.Println("Error 2")
		return nil, getErr
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println("Error 3")
		return nil, readErr
	}

	repos := make([]Repo, 0)
	jsonErr := json.Unmarshal(body, &repos)
	if jsonErr != nil {
		fmt.Println("Error 4")
		return nil, jsonErr
	}
	return repos, nil
}

func main() {

	// todo: 1 - figure out how many pages until error
	// todo: 2 - add all results to single list
	// todo: 3 - persist result in FS
	// todo: 4 - make diff with previous FS and now

	repos, err := extractRepos(5)
	if err == nil {
		for _, r := range repos {
			fmt.Println(r.Name)
		}
	} else {
		fmt.Println("Error")
		fmt.Println(err)
	}

}
