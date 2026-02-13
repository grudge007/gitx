package main

import (
	"fmt"
	"gitx/initz"
	"gitx/pushz"

	// "gitx/pushz"
	"gitx/runx"
	"os"
	"strings"
)

func main() {
	var action string
	// var command []string

	if len(os.Args) > 1 {
		action = os.Args[1]

	} else {
		fmt.Printf("run 'gitx init' to init\nrun 'gitx push' to push code\nrun gitx run <command1> <command2> to execute remote commands in project root")
		return
	}

	if action == "init" && len(os.Args) > 3 {
		// fmt.Printf("%v takes  flags", action)
		return
	}
	switch action {
	case "init":
		force := false
		if len(os.Args) > 2 && (os.Args[2] == "--force" || os.Args[2] == "-f") {
			force = true
		}
		initz.InitGitz(force)

	case "push":
		// pushx.PushFileToRemote()
		loadedConfig := initz.NewInventory().LoadGitzConf()
		pushz.PushFilesToRemote(loadedConfig)

	case "run":
		if len(os.Args) < 3 {
			fmt.Println("No Remote Commands Specified..")
			return
		}
		remoteCommand := strings.Join(os.Args[2:], " ")
		runx.RunCommand(remoteCommand)

	default:
		fmt.Print("Coming Soon")

	}

}
