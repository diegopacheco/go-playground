package main

import (
	"context"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/sethvargo/go-signalcontext"
)

func main() {
	ctx, cancel := signalcontext.OnInterrupt()
	defer cancel()

	println("Running on http://127.0.0.1:8080/ CTRL+C to stop it.")

	s := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for CTRL+C
	<-ctx.Done()

	// Stop the server
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(shutdownCtx); err != nil {
		log.Fatal(err)
	}
}
