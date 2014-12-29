package main

import "fmt"
import "os"
import "os/exec"
import "path"
import "github.com/spf13/viper"

const utility = "rsync"

func init() {
	loadConfig()
}

func main() {
	options := viper.GetStringSlice("options")
	directoryOptions := viper.GetStringMapString("directories")
	directories := []string{
		path.Join(getCurrentDirectory(), directoryOptions["from"]),
		path.Join(getCurrentDirectory(), directoryOptions["to"]),
	}
	sync(append(options, directories...))
}

func sync(options []string) {
	out, error := exec.Command(utility, options...).Output()
	if error != nil {
		fmt.Println("An error occurred while syncing!")
		fmt.Println("%s", error)
	}
	fmt.Printf("%s", out)
}

func loadConfig() {
	viper.SetConfigName("gorsync")
	viper.AddConfigPath(getCurrentDirectory())
	viper.ReadInConfig()
}

func getCurrentDirectory() string {
	pwd, error := os.Getwd()
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	return pwd
}
