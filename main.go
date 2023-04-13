package main

import (
	"fmt"
	"log"
)

func main() {
	if config, err := utils.readConfigFile("/path/to/config.conf"); err == nil {
		fmt.Println("Passes, version:" + config.AppVersion + ", database:" + config.DBVersion)
	} else {
		log.Fatalf("Failed to read config file: %v", err)
	}
	fmt.Println("Passes")
}
