package main

type IntOrString[T int | string] = T

func main() {
	var num IntOrString[int] = 10
	var str IntOrString[string] = "Hello, Go 1.24!"
	println(num)
	println(str)
}
