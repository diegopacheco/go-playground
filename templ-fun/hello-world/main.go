package main

import (
	"context"
	"fmt"
	"net/http"
)

func handleHttp(w http.ResponseWriter, r *http.Request) {
	component := hello("John Doe")
	component.Render(context.Background(), w)
}

func main() {
	http.HandleFunc("/", handleHttp)
	fmt.Print("running on http://127.0.0.1:8080/")
	err := http.ListenAndServe(":8080", nil)
	_ = err
}
