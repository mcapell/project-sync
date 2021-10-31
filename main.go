package main

import (
	"fmt"
)

func main() {
	fmt.Println("Checking local repositories:")
	config, err := GetConfig()
	if err != nil {
		fmt.Println("Unable to get config:", err)
	}
	if err := SyncProjects(config); err != nil {
		fmt.Println("Unable to sync projects:", err)
	}
}
