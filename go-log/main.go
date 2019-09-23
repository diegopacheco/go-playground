package main

import (
  "os"
	"github.com/phuslu/log"
)

func main() {
	log.Debug().Str("foo", "bar").Msg("hello world")

  if log.IsTerminal(os.Stderr.Fd()) {
		log.DefaultLogger = log.Logger{
			Caller: 1,
			Writer: &log.ConsoleWriter{ANSIColor: true},
		}
	}
	log.Info().Str("foo", "bar").Msg("hello world")
}

// Output: {"time":"2019-07-10T16:00:19.092Z","level":"debug","foo":"bar","message":"hello world"}
