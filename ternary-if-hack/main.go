package main

import (
	"fmt"
	"reflect"
)

// ToString converts an interface{} to string
func ToString(o interface{}) string {
	return fmt.Sprintf("%s", o)
}

// inlineIF is a hack to have ternary ifs in go - slower than if-else for sure :D
func inlineIF(condition bool, trueBranch interface{}, falseBranch interface{}) interface{} {
	if condition {
		return trueBranch
	}
	return falseBranch
}

// Invoke calls interface{} by reflection
func Invoke(fn interface{}) {
	reflect.ValueOf(fn).Call(nil)
}

// PrintMethodSet prints all methods from interface{}
func PrintMethodSet(val interface{}) {
	t := reflect.TypeOf(val)
	fmt.Printf("Number of methods: %d\n", t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("Method %s\n", m)
		fmt.Printf("\tName: %s\n", m.Name)
		fmt.Printf("\tPackage path: %s\n", m.PkgPath)
	}
}

func major() {
	fmt.Print("You are major. Awesome. ")
}

func minor() {
	fmt.Print("You are minor. Go watch Peppa Pig. ")
}

var age = 21

func main() {
	result := inlineIF(age >= 18, major, minor)
	fmt.Println(ToString(result))
	PrintMethodSet(result)
	Invoke(result)
}
