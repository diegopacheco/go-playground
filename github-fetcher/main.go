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

func extractRepos(page int, orgname string) ([]Repo, error) {
	url := "https://api.github.com/orgs/" + orgname + "/repos?page=" + strconv.Itoa(page)
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

func getAllRepos(orgname string) []Repo {
	var allRepos []Repo
	pagination := 1
	for true {
		repos, _ := extractRepos(pagination, orgname)
		if len(repos) == 0 {
			break
		}
		for _, r := range repos {
			allRepos = append(allRepos, r)
		}
		pagination = pagination + 1
	}
	return allRepos
}

func main() {
	// todo: 3 - persist result in FS
	// todo: 4 - make diff with previous FS and now
	allRepos := getAllRepos("Netflix")
	for _, o := range allRepos {
		fmt.Println(o.Name)
	}

}
