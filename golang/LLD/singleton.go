package main

import (
	"fmt"
	"sync"
)

// The Singleton Pattern is a design pattern that ensures:
// Only one instance of a particular class exists in your application —
// and that instance is globally accessible.
// Think of it like a single manager in an office:
// Only one person should hold that position.
// Anyone needing a decision goes to that one manager.
// You can’t create more than one manager.
// used for global variables
type Config struct {
	AppName string
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{AppName: "MyApp"}
	})
	return instance
}

// Example of singleton pattern in golang
func singletonExample() {
	cfg1 := GetConfig()
	cfg2 := GetConfig()
	fmt.Println(cfg1 == cfg2) // true → same instance
}
