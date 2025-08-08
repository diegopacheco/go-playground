package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExShellCmd() {
	fmt.Println("ex_shell_cmd")
	out, err := exec.Command("sh", "-c", "echo go && uname -s").Output()
	if err != nil {
		fmt.Println("cmd error:", err)
		return
	}
	s := string(out)
	lines := strings.Split(strings.TrimSpace(s), "\n")
	for _, l := range lines {
		fmt.Println(l)
	}
}
