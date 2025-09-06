package engine

import (
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

func Run(files []string, rules []Rule) ([]Issue, error) {
	var issues []Issue
	var wg sync.WaitGroup
	issueChan := make(chan []Issue)

	for _, f := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			src, err := os.ReadFile(file)
			if err != nil {
				return
			}
			var fileIssues []Issue
			for _, r := range rules {
				res, err := r.Apply(file, src)
				if err != nil {
					continue
				}
				fileIssues = append(fileIssues, res...)
			}
			issueChan <- fileIssues
		}(f)
	}

	go func() {
		wg.Wait()
		close(issueChan)
	}()

	for fileIssues := range issueChan {
		issues = append(issues, fileIssues...)
	}

	return issues, nil
}

func CollectGoFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err == nil && !info.IsDir() && filepath.Ext(path) == ".go" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

