package main

import (
	"os"
	"path/filepath"
)

func findNotes() string {
	gopath := os.Getenv("GOPATH")
	for _, path := range filepath.SplitList(gopath) {
		test := filepath.Join(path, "src", "notes")
		s, err := os.Stat(test)
		if err != nil {
			continue
		}
		if s.IsDir() {
			return test
		}
	}
	return ""
}
