package main

import (
	"fmt"
	"gitx/initx"
	"os"
)

func main() {
	var action string

	if len(os.Args) > 1 {
		action = os.Args[1]

	} else {
		fmt.Printf("run 'gitx init' to init\nrun 'gitx push' to push code\nrun gitx run <command1> <command2> to execute remote commands in project root")
		return
	}

	if action == "init" && len(os.Args) > 2 {
		fmt.Printf("%v takes zero flags", action)
		return
	}
	switch action {
	case "init":
		initx.InitializeConfig()
	case "push":

	default:
		fmt.Print("Coming Soon")

	}

}
