package main

import (
	"fmt"
	"log"

	"github.com/lxzan/gws"
	kcp "github.com/xtaci/kcp-go"
)

func main() {
	fmt.Print("WS client running on 6666")
	conn, err := kcp.Dial("127.0.0.1:6666")
	if err != nil {
		log.Println(err.Error())
		return
	}
	app, _, err := gws.NewClientFromConn(&gws.BuiltinEventHandler{}, nil, conn)
	if err != nil {
		log.Println(err.Error())
		return
	}
	app.ReadLoop()
}
