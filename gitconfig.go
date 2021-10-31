package main

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"
)

func GitExcludes() []string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user", err)
	}
	file, _ := ioutil.ReadFile(filepath.Join(usr.HomeDir, ".gitignore_global"))
	return strings.Split(string(file), "\n")
}
