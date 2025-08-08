package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID     uint32 `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

func ExJSONSerde() {
	fmt.Println("ex_json_serde")

	u := User{ID: 1, Name: "neo", Active: true}
	s, err := json.Marshal(u)
	if err != nil {
		fmt.Println("marshal error:", err)
		return
	}
	fmt.Println("json:", string(s))

	var d User
	if err := json.Unmarshal(s, &d); err != nil {
		fmt.Println("unmarshal error:", err)
		return
	}
	fmt.Println("struct:", d)
}
