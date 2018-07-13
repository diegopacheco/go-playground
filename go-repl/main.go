package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func printRepl() {
	fmt.Print("go-repl> ")
}

func printInvalidCmd(text string) {
	t := strings.TrimSuffix(text, "\n")
	if t != "" {
		fmt.Println("go-repl> unknow command " + t)
	}
}

func get(r *bufio.Reader) string {
	t, _ := r.ReadString('\n')
	return strings.TrimSpace(t)
}

func shouldContinue(text string) bool {
	if strings.EqualFold("exit", text) {
		return false
	}
	return true
}

func help() {
	fmt.Println("go-repl> Welcome to Go Repl! ")
	fmt.Println("go-repl> Wrote by Diego Pacheco - 2018 ")
	fmt.Println("go-repl> This Are the Avaliable commands: ")
	fmt.Println("go-repl> help - Show you the Help")
	fmt.Println("go-repl> cls  - Clear the Terminal Screen ")
	fmt.Println("go-repl> exit - Exits the Go REPL ")
	fmt.Println("go-repl> ")
}

func cls() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	commands := map[string]interface{}{
		"help": help,
		"cls":  cls,
	}
	reader := bufio.NewReader(os.Stdin)
	help()
	printRepl()
	text := get(reader)
	for ; shouldContinue(text); text = get(reader) {
		if value, exists := commands[text]; exists {
			value.(func())()
		} else {
			printInvalidCmd(text)
		}
		printRepl()
	}
	fmt.Println("Bye!")

}
