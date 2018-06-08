package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Repo represents github repository from org
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

func persistInDisk(path string, v interface{}) error {
	if v == nil {
		return nil
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	binary, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, bytes.NewReader(binary))
	f.Close()
	return err
}

func loadFromDisk(path string, v *[]Repo) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	f.Close()
	json.NewDecoder(f).Decode(&v)
	return nil
}

func diff(slice1 []Repo, slice2 []Repo) []Repo {
	diffRepo := []Repo{}
	m := map[Repo]int{}
	for _, s1Val := range slice1 {
		m[s1Val] = 1
	}
	for _, s2Val := range slice2 {
		m[s2Val] = m[s2Val] + 1
	}
	for mKey, mVal := range m {
		if mVal == 1 {
			diffRepo = append(diffRepo, mKey)
		}
	}
	return diffRepo
}

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("ERROR! You need provide the name of the organization. ")
		return
	}

	fmt.Println("Loading previous JSON from disk: ")
	var allReposFromDisk []Repo
	if loadFromDisk("/home/diego/github.fetcher/"+args[1]+".json", &allReposFromDisk) == nil {
		fmt.Print("No Previous JSON in DISK.\n")
	} else {
		fmt.Println("JSON from disk:")
		fmt.Println(allReposFromDisk)
	}

	fmt.Println("Fetching all repos for: " + args[1])
	allRepos := getAllRepos(args[1])
	for _, o := range allRepos {
		fmt.Println(o.Name)
	}

	if persistInDisk("/home/diego/github.fetcher/"+args[1]+".json", allRepos) != nil {
		fmt.Println("Json perssited in DISK!")
	} else {
		fmt.Println("JSON persisted in disk")
	}

	diffRepo := diff(allReposFromDisk, allRepos)
	fmt.Println("\n\n**** NEW REPOS **** ")
	fmt.Println(diffRepo)
}
