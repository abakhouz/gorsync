package main

import (
	"path/filepath"
	"testing"
)

func TestGetCurrentDirectory(t *testing.T) {
	workingDirectory, _ := filepath.Abs("./")
	value, _ := getCurrentDirectory()
	if value != workingDirectory {
		t.Error(
			"getCurrentDirectory does not return current working directory:",
			workingDirectory,
		)
	}
}
