package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"
)

const DEFAULT_PATH = "~/.config/project_sync.yaml"

type Project struct {
	Name       string
	Repository string `yaml:"repo"`
	Path       string
}

type Config struct {
	Projects []Project
}

func resolveUserPath(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr.HomeDir, path[2:]), nil
}

func resolveConfigPaths(config *Config) {
	for i, project := range config.Projects {
		resolved, err := resolveUserPath(project.Path)
		if err != nil {
			fmt.Println("Unable to resolve user path:", err)
		}
		config.Projects[i].Path = resolved
	}
}

func GetConfig() (Config, error) {
	config := Config{}
	path, err := resolveUserPath(DEFAULT_PATH)
	if err != nil {
		fmt.Println("Unable to get configuration path", err)
		return config, err
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Unable to read contents", err)
		return config, err
	}

	yaml.Unmarshal(content, &config)
	resolveConfigPaths(&config)
	return config, nil
}
