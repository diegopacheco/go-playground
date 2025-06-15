package main

import (
	"fmt"
	"sync"
	"time"
)

type Config struct {
	AppName string
}

var (
	singleton *Config
	once      sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		fmt.Println("Initializing config...")
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		singleton = &Config{AppName: fmt.Sprintf("MyGoApp - %s", currentTime)}
	})
	return singleton
}

func main() {
	cfg1 := GetConfig()
	time.Sleep(2 * time.Second)
	cfg2 := GetConfig()
	time.Sleep(2 * time.Second)
	cfg2 = GetConfig()
	cfg2 = GetConfig()
	cfg2 = GetConfig()
	cfg2 = GetConfig()
	fmt.Printf("Config 1: %v\n", cfg1)
	fmt.Printf("Config 2: %v\n", cfg2)
	fmt.Println("Same instance?", cfg1 == cfg2)
}
