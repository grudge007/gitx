package main

import (
	"fmt"
	"gitx/initz"
	"gitx/pushz"
	"gitx/runz"

	"os"
	"strings"
)

const Version = "0.1.1"

func main() {
	var action string

	if len(os.Args) > 1 {
		action = os.Args[1]
	} else {
		fmt.Println("Usage:")
		fmt.Println("  gitx init [--force|-f]")
		fmt.Println("  gitx push")
		fmt.Println("  gitx run <command>")
		fmt.Println("  gitx --version")
		return
	}

	if action == "init" && len(os.Args) > 3 {
		fmt.Println("Error: too many arguments for 'gitx init'.")
		fmt.Println("Usage: gitx init [--force|-f]")
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
		loadedConfig := initz.NewInventory().LoadGitzConf()
		pushz.PushFilesToRemote(loadedConfig)

	case "run":
		if len(os.Args) < 3 {
			fmt.Println("Error: no remote command specified.")
			fmt.Println("Usage: gitx run <command>")
			return
		}
		remoteCommand := strings.Join(os.Args[2:], " ")
		loadedConfig := initz.NewInventory().LoadGitzConf()
		runz.RunCommand(loadedConfig, remoteCommand)

	case "-v", "--version":
		fmt.Printf("gitx version %s\n", Version)
		return

	default:
		fmt.Printf("Error: unknown command '%s'.\n", action)
		fmt.Println("Usage:")
		fmt.Println("  gitx init [--force|-f]")
		fmt.Println("  gitx push")
		fmt.Println("  gitx run <command>")
		fmt.Println("  gitx --version")
	}
}
