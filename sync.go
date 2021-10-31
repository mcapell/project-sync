package main

import (
	"fmt"
	"sync"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
)

func SyncProjects(config Config) error {
	var wg sync.WaitGroup
	result := make(chan string)
	for _, conf := range config.Projects {
		wg.Add(1)
		go func(project Project) {
			defer wg.Done()
			syncProject(project, result)
		}(conf)
	}
	go func() {
		wg.Wait()
		close(result)
	}()

	for res := range result {
		fmt.Println(res)
	}
	return nil
}

func syncProject(project Project, out chan string) {
	repo, err := git.PlainOpen(project.Path)
	if err != nil {
		out <- "Found some error with repo:" + project.Name
		return
	}
	wt, err := repo.Worktree()
	if err != nil {
		out <- "Error in worktree:" + project.Name
		return
	}
	for _, pattern := range(GitExcludes()) {
		wt.Excludes = append(wt.Excludes, gitignore.ParsePattern(pattern, nil))
	}
	status, err := wt.Status()
	if err != nil {
		out <- "❌ (Error retrieving the status):" + project.Name
		return
	}
	if !status.IsClean() {
		out <- "❌ (Uncommit work): " + project.Name
		return
	}
	out <- "✅: " + project.Name
	return

}
