package main

import "github.com/cheynewallace/tabby"

func main() {
	t := tabby.New()
	t.AddHeader("NAME", "TITLE", "DEPARTMENT")
	t.AddLine("John Smith", "Developer", "Engineering")
	t.Print()

	print("\n")
	print("\n")

	t = tabby.New()
	t.AddLine("Info:", "WEB", "Success 200")
	t.AddLine("Info:", "API", "Success 201")
	t.AddLine("Error:", "DATABASE", "Connection Established")
	t.Print()
}
