package main

import "github.com/tidwall/evio"

//
// go run telnet-evio.go
//
// new terminal do $ echo "diego" | telnet localhost 5000
//
func main() {
	print("Goto telnet localhost 5000... ")
	var events evio.Events
	events.Data = func(id int, in []byte) (out []byte, action evio.Action) {
	  out = in
	  print(in) 
	  print("\n\r")
	  return
	}
	if err := evio.Serve(events, "tcp://localhost:5000"); err != nil {
	  panic(err.Error())
	}
}
