package main

import "fmt"

type Greeter interface {
	Greet() string
}

type Person struct{ Name string }

func (p Person) Greet() string { return "hello " + p.Name }

func greetAll(items []Greeter) {
	for _, it := range items {
		fmt.Println(it.Greet())
	}
}

func ExTraits() {
	fmt.Println("ex_traits")
	people := []Greeter{Person{"neo"}, Person{"trinity"}}
	greetAll(people)
}
