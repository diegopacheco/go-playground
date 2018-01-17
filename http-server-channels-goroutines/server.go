package main

import (
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"
)

var counter *big.Int

func requestHandler1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func requestHandler2(w http.ResponseWriter, r *http.Request) {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() { c1 <- "ok 1" }()
	go func() { c2 <- "ok 2" }()
	go func() { c3 <- "ok 3" }()

	r1, r2, r3 := <-c1, <-c2, <-c3
	counter = big.NewInt(0).Add(counter, big.NewInt(1))
	data := time.Now().Format(time.RFC850) + " [" + counter.String() + "]" + " /get -> All results: [c1:" + r1 + " c2:" + r2 + " c3: " + r3 + "]"
	io.WriteString(w, data)
	fmt.Println("Data: ", data)
}

func main() {
	counter = big.NewInt(0)
	fmt.Println("Started GO Server on 8080... ")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() == "/hello" {
		requestHandler1(w, r)
	} else if r.URL.String() == "/get" {
		requestHandler2(w, r)
	} else {
		io.WriteString(w, "My server: "+r.URL.String())
	}
}
