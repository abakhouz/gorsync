package main

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	olog *log.Logger
	elog *log.Logger
)

const configFileName = "gorsync"

func init() {
	olog = log.New(os.Stdout, "", 0)
	elog = log.New(os.Stderr, "", 0)
}

func main() {
	loadConfig()
	r := new(rsync)
	r.sync(r.generateOptions())
	olog.Println("Sycing succesful!")
}

func loadConfig() {
	viper.SetConfigName(configFileName)
	configDirectory, _ := getCurrentDirectory()
	viper.AddConfigPath(configDirectory)
	error := viper.ReadInConfig()
	if error != nil {
		elog.Fatal("No ./gorsync.yml|json|toml config file found!")
	}
}
