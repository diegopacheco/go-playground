package main

import (
	"github.com/phuslu/log"
)

func main() {
	log.Debug().Str("foo", "bar").Msg("hello world")
}

// Output: {"time":"2019-07-10T16:00:19.092Z","level":"debug","foo":"bar","message":"hello world"}
