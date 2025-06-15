package main

import "fmt"

type Plugin interface {
	Run(input string) string
}

type UppercasePlugin struct{}

func (u UppercasePlugin) Run(input string) string {
	return fmt.Sprintf("Uppercase: %s", upper(input))
}

type ReversePlugin struct{}

func (r ReversePlugin) Run(input string) string {
	return fmt.Sprintf("Reversed: %s", reverse(input))
}

func upper(s string) string {
	return string([]rune(s))
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func executePlugins(input string, plugins []Plugin) {
	for _, p := range plugins {
		fmt.Println(p.Run(input))
	}
}

func main() {
	plugins := []Plugin{
		UppercasePlugin{},
		ReversePlugin{},
	}
	executePlugins("GoLang", plugins)
}
