package main

import "fmt"

// Repo represents github repository from org
type Repo struct {
	Name string `json:"full_name"`
}

func diff(slice1 []Repo, slice2 []Repo) []string {
	diffRepo := make([]string, 0)
	m := map[string]int{}
	for _, s1Val := range slice1 {
		m[s1Val.Name] = 1
	}
	for _, s2Val := range slice2 {
		_, ok := m[s2Val.Name]
		if ok {
			m[s2Val.Name] = m[s2Val.Name] + 1
		} else {
			m[s2Val.Name] = 1
		}
	}
	for mKey, mVal := range m {
		if mVal == 1 {
			diffRepo = append(diffRepo, mKey)
		}
	}
	return diffRepo
}

func main() {
	repo1 := make([]Repo, 0)
	repo2 := make([]Repo, 0)

	repo1[0] = Repo{}
	repo1[1] = Repo{}
	repo1[2] = Repo{}

	repo1[0].Name = "X"
	repo1[1].Name = "Y"
	repo1[2].Name = "Z"

	repo2[0] = Repo{}
	repo2[1] = Repo{}
	repo2[2] = Repo{}
	repo2[3] = Repo{}

	repo2[0].Name = "X"
	repo2[1].Name = "Y"
	repo2[2].Name = "Z"
	repo2[3].Name = "W"

	diffRepo := diff(repo1, repo2)
	fmt.Println(diffRepo)

}
