package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func printStructFields(data interface{}) {
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s, Type: %s, Tag: %s\n",
			t.Field(i).Name,
			t.Field(i).Type,
			t.Field(i).Tag.Get("json"))
	}
}

func main() {
	u := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	printStructFields(u)
}
