package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	NumberOfNodes int
	DomainOrip    []string
	ProjectPath   string
}

func main() {

	var ConfigDir = ".gitx"
	var ConfigFIle = ".gitx/gitx.conf"
	var action string
	// var command []string
	var numberOfNodes int
	var projectPath string

	var nodes []string

	if len(os.Args) > 1 {
		action = os.Args[1]
		// command = os.Args[2:]

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
		var domainOrip string
		var userConfig Config

		fmt.Print("How Many Nodes You Have: ")
		fmt.Scan(&numberOfNodes)

		for i := 1; i <= numberOfNodes; i++ {
			fmt.Printf("domain name or ip address of node %d: ", i)
			fmt.Scan(&domainOrip)
			nodes = append(nodes, domainOrip)
		}

		fmt.Printf("Project Directory From '/': ")
		fmt.Scan(&projectPath)
		os.MkdirAll(ConfigDir, 0755)
		file, _ := os.Create(ConfigFIle)
		defer file.Close()

		userConfig = Config{
			NumberOfNodes: numberOfNodes,
			DomainOrip:    nodes,
			ProjectPath:   projectPath,
		}
		fmt.Println("userConfig: ", userConfig)

		userConfigJson, _ := json.MarshalIndent(userConfig, "", "  ")
		fmt.Println("userConfigJson: ", string(userConfigJson))
		file.Write(userConfigJson)
	default:
		fmt.Print("Coming Soon")

	}

}
