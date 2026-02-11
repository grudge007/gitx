package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
)

type Node struct {
	IP   string `yaml:"ip"`
	User string `yaml:"user"`
	Path string `yaml:"path"`
}

type Inventory struct {
	ProjectName string `yaml:"project_name"`
	DefaultUser string `yaml:"default_user"`
	DefaultPath string `yaml:"default_path"`
	Nodes       []Node
}

func main() {
	gitzConfFile := workingDirAndConfigFile()
	fmt.Println(gitzConfFile)
	isExist := checkIsGitzExist(gitzConfFile)
	if !isExist {
		return
	}

	inventoryFile := Inventory{
		ProjectName: "MyProject",
		DefaultUser: "root",
		DefaultPath: "/opt/MyProject",
		Nodes: []Node{
			{
				IP:   "192.168.100.10",
				User: "root",
				Path: "/root/mypath",
			},
		},
	}
	inventoryYaml := convertToYaml(inventoryFile)
	saveToGitzConf(inventoryYaml, gitzConfFile)

}

func workingDirAndConfigFile() string {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(".gitz", 0766)
	gitzconfigFile := filepath.Join(workingDir, ".gitz/gitz.yaml")
	return gitzconfigFile
}

func checkIsGitzExist(gitzConfFile string) bool {
	_, err := os.Stat(gitzConfFile)
	if err == nil {
		if len(os.Args) >= 2 {
			if os.Args[1] == "-f" {
			} else {
				fmt.Println("an existing configuration found")
				fmt.Println("run with -f to override")
				return false
			}
		} else {
			fmt.Println("an existing configuration found")
			fmt.Println("run with -f to override")
			return false
		}
	}
	return true
}

func convertToYaml(inventoryFile Inventory) []byte {
	inventoryYaml, err := yaml.Marshal(&inventoryFile)
	if err != nil {
		log.Fatal(err)
	}
	return inventoryYaml
}

func saveToGitzConf(inventoryYaml []byte, gitzConfFile string) {
	err := os.WriteFile(gitzConfFile, inventoryYaml, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(inventoryYaml))
}
