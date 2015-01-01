package main

import (
	"bytes"
	"github.com/spf13/viper"
	"log"
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

func TestSyncOutputsOutputIfAny(t *testing.T) {
	setupConfig()
	oldOlog := olog
	buffer := new(bytes.Buffer)
	olog = log.New(buffer, "", 0)

	viper.Set("options", []string{"-av"})
	r := r()
	r.sync(r.generateOptions())

	olog = oldOlog

	if len(buffer.String()) == 0 {
		t.Error("rsync output not outputted to stdout")
	}

	cleanTestDir()
}

func TestGenerateOptions(t *testing.T) {
	expectedOptions := strings.Join([]string{
		"-au -v",
		path.Join("test", "data_dir"),
		path.Join("test", "sync"),
	}, " ")

	viper.Set("options", []string{"-au", "-v"})
	viper.Set("directories", map[string]string{
		"from": "test/data_dir",
		"to":   "test/sync",
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
