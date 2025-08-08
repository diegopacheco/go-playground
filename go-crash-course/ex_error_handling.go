package main

import (
	"errors"
	"fmt"
	"os"
)

func ExErrorHandling() {
	fmt.Println("ex_error_handling")

	content, err := readFileToString("nonexistent.txt")
	if err != nil {
		fmt.Println("failed:", err)
	} else {
		fmt.Println("file:", len(content))
	}

	if r, err := mayReturn(1); err == nil {
		fmt.Println("ok:", r)
	} else {
		fmt.Println("err:", err)
	}
}

func readFileToString(p string) (string, error) {
	b, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func mayReturn(x int) (int, error) {
	if x > 0 {
		return x, nil
	}
	return 0, errors.New("x must be > 0")
}
