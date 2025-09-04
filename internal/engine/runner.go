package engine

import (
	"io/fs"
	"os"
	"path/filepath"
)

func Run(files []string, rules []Rule) ([]Issue, error) {
	var issues []Issue
	for _, f := range files {
		src, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}
		// TODO: we need to make this concurrent
		for _, r := range rules {
			res, err := r.Apply(f, src)
			if err != nil {
				return nil, err
			}
			issues = append(issues, res...)
		}
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
