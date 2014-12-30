package main

import (
	"bytes"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"path"
	"strings"
)

const utility = "rsync"

type rsync struct{}

func (r *rsync) sync(options []string) {
	cmd := exec.Command(utility, options...)
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if error := cmd.Run(); error != nil {
		elog.Fatal(stderr.String() + "\n" + error.Error())
	}

	stdoutString := stdout.String()
	trmString := strings.TrimSpace(stdoutString)
	if len(trmString) > 0 {
		olog.Println(trmString)
	}
}

func (r *rsync) generateOptions() []string {
	options := viper.GetStringSlice("options")
	directoryOptions := viper.GetStringMapString("directories")
	currentDirectory, error := getCurrentDirectory()

	if error != nil {
		elog.Fatal(error.Error())
	}

	directories := []string{
		path.Join(currentDirectory, directoryOptions["from"]),
		path.Join(currentDirectory, directoryOptions["to"]),
	}
	return append(options, directories...)
}
