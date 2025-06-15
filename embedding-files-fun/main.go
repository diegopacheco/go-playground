package main

import (
	"embed"
	"fmt"
)

//go:embed hello.txt
var content embed.FS

func main() {
	data, err := content.ReadFile("hello.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
