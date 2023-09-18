package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/a-h/templ-examples/hello-world/src/db"
	"github.com/a-h/templ-examples/hello-world/src/handlers"
	"github.com/a-h/templ-examples/hello-world/src/services"
	"github.com/a-h/templ-examples/hello-world/src/session"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	s, _ := db.NewCountStore()
	cs := services.NewCount(log, s)
	h := handlers.New(log, cs)
	sh := session.NewMiddleware(h, session.WithSecure(false))

	server := &http.Server{
		Addr:         "localhost:9000",
		Handler:      sh,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	fmt.Printf("Listening on http://%v\n", server.Addr)
	server.ListenAndServe()
}
