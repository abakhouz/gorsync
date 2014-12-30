package main

import (
	"bytes"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

const utility = "rsync"

var (
	olog *log.Logger
	elog *log.Logger
)

func init() {
	loadConfig()
	olog = log.New(os.Stdout, "", 0)
	elog = log.New(os.Stderr, "", 0)
}

func main() {
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
	sync(append(options, directories...))
	olog.Println("Sycing succesful!")
}

func sync(options []string) {
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

func loadConfig() {
	viper.SetConfigName("gorsync")
	configDirectory, _ := getCurrentDirectory()
	viper.AddConfigPath(configDirectory)
	viper.ReadInConfig()
}

func getCurrentDirectory() (string, error) {
	return os.Getwd()
}
