package initx

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

type Config struct {
	NumberOfNodes   int
	InventoryStruct Inventory
}

type Inventory struct {
	UserName       string
	NodeIpOrDomain string
	ProjectPath    string
}

func InitializeConfig() {
	var ConfigDir = ".gitx"
	var ConfigFIle = ".gitx/gitx.conf"
	var numberOfNodes string
	var projectPath string
	var domainOrip string
	var userConfig Config
	var isOverride string
	var userName string
	var inventory = Inventory{}

	_, err := os.Stat(ConfigDir)
	if err == nil {
		_, err = os.Stat(ConfigFIle)
		if err == nil {
			fmt.Println("existing configuration found")
			fmt.Print("Enter 'Y' to override: ")
			fmt.Scan(&isOverride)
			if isOverride != "y" && isOverride != "Y" {
				return
			}
		}

	} else {
		err = os.MkdirAll(ConfigDir, 0755)
		if err != nil {
			fmt.Println("Failed to create config file")
			return
		}
	}

	fmt.Print("How Many Nodes You Have: ")
	fmt.Scan(&numberOfNodes)
	numberOfNode, err := strconv.Atoi(numberOfNodes)
	if err != nil {
		return
	}

	for i := 1; i <= numberOfNode; i++ {
		fmt.Printf("domain name or ip address of node %d: ", i)
		fmt.Scan(&domainOrip)
		fmt.Printf("Username of %d: ", i)
		fmt.Scan(&userName)
		if len(domainOrip) > 50 {
			fmt.Printf("%v characters?? looks abnormal!!\n", len(domainOrip))
			return
		}

		inventory.NodeIpOrDomain = domainOrip
		inventory.UserName = userName

	}

	fmt.Printf("Project Directory From '~' (eg: ~/my/path): ")
	fmt.Scan(&projectPath)
	projectPath = strings.ReplaceAll(projectPath, "~", "")
	projectPath = setProjectPath(projectPath)
	inventory.ProjectPath = projectPath

	_, err = os.Stat(projectPath)
	if err != nil {
		fmt.Printf("%v path does not exists quitting..\n", projectPath)
		return
	}
	inventoryConfig := Inventory{
		UserName:       inventory.UserName,
		NodeIpOrDomain: inventory.NodeIpOrDomain,
		ProjectPath:    inventory.ProjectPath,
	}

	userConfig = Config{
		NumberOfNodes:   numberOfNode,
		InventoryStruct: inventoryConfig,
	}
	fmt.Println("userConfig: ", userConfig)

	userConfigJson, err := json.MarshalIndent(userConfig, "", "  ")
	if err != nil {
		fmt.Println("parse error")
		return
	}

	fmt.Println("userConfigJson: ", string(userConfigJson))
	err = os.WriteFile(ConfigFIle, userConfigJson, 0660)

	if err != nil {
		fmt.Printf("Failed to write input data to %v\n", ConfigFIle)
		return
	}
}

func setProjectPath(projectPath string) string {
	currentUserData, err := user.Current()
	if err != nil {
		return ""
	}
	userName := currentUserData.Username
	projectPath = filepath.Join("/home/", userName, projectPath)
	return projectPath

}
