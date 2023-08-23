package main

import (
	"log/slog"
	"os"
)

func main() {
	slog.Info("hello, world")
	slog.Info("hello, world", "javahome", os.Getenv("JAVA_HOME"))
}
