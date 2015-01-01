package main

import (
	"bytes"
	"github.com/spf13/viper"
	"io"
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
	re, w, _ := os.Pipe()
	oldOlog := olog
	olog = log.New(w, "", 0)

	viper.Set("options", []string{"-av"})
	r := r()
	r.sync(r.generateOptions())

	outC := make(chan string)
	go func() {
		buf := new(bytes.Buffer)
		io.Copy(buf, re)
		outC <- buf.String()
	}()

	w.Close()
	out := <-outC
	olog = oldOlog

	if len(out) == 0 {
		t.Error("rsync output not outputted to stdout")
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
