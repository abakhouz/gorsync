package main

import "fmt"
import "os"
import "os/exec"
import "path"

const utility = "rsync"

func main() {
	options := []string{
		"-a",
		path.Join(getCurrentDirectory(), "a"),
		path.Join(getCurrentDirectory(), "b"),
	}
	sync(options)
}

func sync(options []string) {
	out, error := exec.Command(utility, options...).Output()
	if error != nil {
		fmt.Println("An error occurred while syncing!")
		fmt.Println("%s", error)
	}
	fmt.Printf("%s", out)
}

func getCurrentDirectory() string {
	pwd, error := os.Getwd()
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	return pwd
}
