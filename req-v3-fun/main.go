package main

import (
	"fmt"

	"github.com/imroc/req"
)

type Repo struct {
	Name string `json:"name"`
	Star int    `json:"stargazers_count"`
}

func main() {
	var repos []interface{}
	var username = "diegopacheco"
	resp, _ := req.Get("https://api.github.com/users/" + username + "/repos")
	resp.ToJSON(&repos)

	for _, repo := range repos {
		r := repo.(map[string]interface{})
		fmt.Println(r["name"], r["stargazers_count"].(float64))
	}
}
