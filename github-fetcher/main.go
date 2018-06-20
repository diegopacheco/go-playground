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
		fmt.Println("Error 4 - Often this means we HIT rate limit of Github API")
		return nil, jsonErr
	}
	return repos, nil
}

func getAllRepos(orgname string) []Repo {
	var allRepos []Repo
	pagination := 0
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

func persistInDisk(path string, v []Repo) error {
	if v == nil || len(v) == 0 {
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
	json.NewDecoder(f).Decode(&v)
	f.Close()
	return nil
}

func diff(slice1 []Repo, slice2 []Repo) []string {
	diffRepo := make([]string, 0)
	m := map[string]int{}
	for _, s1Val := range slice1 {
		m[s1Val.Name] = 1
	}
	for _, s2Val := range slice2 {
		m[s2Val.Name] = m[s2Val.Name] + 1
	}
	for mKey, mVal := range m {
		if mVal == 1 {
			diffRepo = append(diffRepo, mKey)
		} else {
			fmt.Printf("mVal is: %d key is: %s \n", mVal, mKey)
		}
	}
	fmt.Print("Diff is: ")
	fmt.Println(diffRepo)
	return diffRepo
}

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println(" _____ _ _   _           _      ______        _   _               ")
		fmt.Println("|  __ (_) | | |         | |     |  ___|      | | | |              ")
		fmt.Println("| |  \\/_| |_| |__  _   _| |__   | |_ ___  ___| |_| |__   ___ _ __ ")
		fmt.Println("| | __| | __| '_ \\| | | | '_ \\  |  _/ _ \\/ __| __| '_ \\ / _ \\ '__|")
		fmt.Println("| |_\\ \\ | |_| | | | |_| | |_) | | ||  __/ (__| |_| | | |  __/ |   ")
		fmt.Println(" \\____/_|\\__|_| |_|\\__,_|_.__/  \\_| \\___|\\___|\\__|_| |_|\\___|_|    ")
		fmt.Println(" ")
		fmt.Println("github-fecther: Fetch all repos from a organization and tell you the new repos! ")
		fmt.Println("USAGE         : ./github-fecther Facebook")
		fmt.Println("BY            : Diego Pacheco - 2018")
		return
	}

	fmt.Println("1. Loading previous JSON from disk: ")
	allReposFromDisk := make([]Repo, 0)
	if loadFromDisk("/home/diego/github.fetcher/"+args[1]+".json", &allReposFromDisk) != nil {
		fmt.Print("No Previous JSON in DISK.\n")
	} else {
		fmt.Println("JSON from disk:")
		for _, o := range allReposFromDisk {
			fmt.Println(o.Name)
		}
	}

	fmt.Println("2. Fetching all repos for: " + args[1])
	allRepos := getAllRepos(args[1])
	for _, o := range allRepos {
		fmt.Println(o.Name)
	}

	if allRepos == nil || len(allRepos) == 0 {
		fmt.Println("There are no repos on the WEB we will not proceed! ")
	} else {

		fmt.Print("Repos from DISK  : ")
		fmt.Println(len(allReposFromDisk))
		fmt.Print("Repos from Github: ")
		fmt.Println(len(allRepos))

		diffRepo := diff(allReposFromDisk, allRepos)
		fmt.Println("\n\n3. **** NEW REPOS **** ")
		fmt.Println(diffRepo)

		if persistInDisk("/home/diego/github.fetcher/"+args[1]+".json", allRepos) != nil {
			fmt.Println("JSON NOT persisted in disk")
		} else {
			fmt.Println("JSON persisted in disk")
		}

	}

}
