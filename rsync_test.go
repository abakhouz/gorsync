package main

import (
	"github.com/spf13/viper"
	"os"
	"path"
	"strings"
	"testing"
)

func setupConfig() {
	viper.Set("options", []string{"-a"})
	viper.Set("directories", map[string]string{
		"from": "test/data_dir",
		"to":   "test/sync",
	})
}

func TestSync(t *testing.T) {
	setupConfig()

	r := r()
	r.sync(r.generateOptions())
	currentPath, _ := getCurrentDirectory()
	resultDirectory := path.Join(currentPath, "test", "sync", "data_dir")
	if _, err := os.Stat(resultDirectory); os.IsNotExist(err) {
		t.Error("rsync command failed")
	}

	cleanTestDir()
}

func TestGenerateOptions(t *testing.T) {
	currentPath, _ := getCurrentDirectory()
	expectedOptions := strings.Join([]string{
		"-au -v",
		path.Join(currentPath, "test", "data_dir"),
		path.Join(currentPath, "test", "sync", "data_dir"),
	}, " ")

	viper.Set("options", []string{"-au", "-v"})
	viper.Set("directories", map[string]string{
		"from": "test/data_dir",
		"to":   "test/sync/data_dir",
	})
	returnedOptions := strings.Join(r().generateOptions(), " ")

	if returnedOptions != expectedOptions {
		t.Error(
			"rsync.generateOptions()",
			"Expected:",
			expectedOptions,
			"Got:",
			returnedOptions,
		)
	}
}

func r() *rsync {
	return new(rsync)
}

func cleanTestDir() {
	currentPath, _ := getCurrentDirectory()
	resultDirectory := path.Join(currentPath, "test", "sync")
	os.RemoveAll(resultDirectory)
}
