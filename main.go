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

func init() {
	olog = log.New(os.Stdout, "", 0)
	elog = log.New(os.Stderr, "", 0)
	loadConfig()
}

func main() {
	r := new(rsync)
	r.sync(r.generateOptions())
	olog.Println("Sycing succesful!")
}

func loadConfig() {
	viper.SetConfigName("gorsync")
	configDirectory, _ := getCurrentDirectory()
	viper.AddConfigPath(configDirectory)
	error := viper.ReadInConfig()
	if error != nil {
		elog.Fatal("No ./gorsync.yml|json|toml config file found!")
	}
}
