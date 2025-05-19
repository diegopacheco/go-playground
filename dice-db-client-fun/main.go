package main

import (
	"fmt"

	"github.com/dicedb/dicedb-go"
	"github.com/dicedb/dicedb-go/wire"
)

func main() {
	client, err := dicedb.NewClient("localhost", 7379)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	key := "k1"
	value := "v1"
	resp := client.Fire(&wire.Command{
		Cmd:  "SET",
		Args: []string{key, value},
	})
	if resp.Status == wire.Status_ERR {
		fmt.Println("error setting key:", resp.Message)
		return
	}
	fmt.Printf("successfully set key %s=%s\n", key, value)

	resp = client.Fire(&wire.Command{
		Cmd:  "GET",
		Args: []string{key},
	})
	if resp.Status == wire.Status_ERR {
		fmt.Println("error getting key:", resp.Message)
		return
	}
	fmt.Printf("successfully got key %s=%s\n", key, resp.GetGETRes().Value)
}
