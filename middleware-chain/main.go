package main

import (
	"fmt"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request URI:", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Auth") != "secret" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func finalHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Request passed all middleware")
}

func main() {
	final := http.HandlerFunc(finalHandler)
	http.Handle("/", loggingMiddleware(authMiddleware(final)))
	http.ListenAndServe(":8080", nil)
}
