package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

type MyHandler struct {
	foobar string
}

func (h *MyHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, world! Requested path is %q. Foobar is %q",
		ctx.Path(), h.foobar)
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
}

func main() {
	myHandler := &MyHandler{
		foobar: "foobar",
	}
	fmt.Print("Running fasthttp on 127.0.0.1:8080 and 127.0.0.1:8081")
	fasthttp.ListenAndServe(":8080", myHandler.HandleFastHTTP)
	fasthttp.ListenAndServe(":8081", fastHTTPHandler)
}
