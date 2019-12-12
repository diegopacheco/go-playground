package main

import (
	"fmt"
	"github.com/golobby/config"
	"github.com/golobby/config/feeder"
)

func main() {
	c, _ := config.New(config.Options{})
	c.Set("name", "John Doe")
	name, _ := c.Get("name")
	fmt.Printf("%v\n", name)

	c, err := config.New(config.Options{
		Feeder: feeder.Map{
			"name":     "Hey You",
			"band":     "Pink Floyd",
			"year":     1979,
			"duration": 4.6,
		},
	})
	if err != nil {
		panic(err)
	}

	result, _ := c.GetString("name")
	fmt.Printf("%v\n", result)
}
