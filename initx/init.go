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
	UserName       []string
	PassWord       []string
	NodeIpOrDomain []string
	ProjectPath    []string
}

var ConfigFile = ".gitx/gitx.conf"

func InitializeConfig() {
	var ConfigDir = ".gitx"
	var numberOfNodes string
	var projectPath []string
	var domainOrip []string
	var userConfig Config
	var isOverride string
	var userName []string
	var password []string
	// var inventory = Inventory{}

	_, err := os.Stat(ConfigDir)
	if err == nil {
		_, err = os.Stat(ConfigFile)
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
		var ipOrDom string
		var pass string
		var user string
		var projectRoot string
		fmt.Printf("domain name or ip address of node %d: ", i)
		fmt.Scan(&ipOrDom)
		fmt.Printf("Username of %d: ", i)
		fmt.Scan(&user)
		fmt.Printf("Password of %d: ", i)
		fmt.Scan(&pass)
		if len(ipOrDom) > 50 {
			fmt.Printf("%v characters?? looks abnormal!!\n", len(domainOrip))
			return
		}
		fmt.Printf("Project Directory From '~' (eg: ~/my/path): ")
		fmt.Scan(&projectRoot)

		domainOrip = append(domainOrip, fmt.Sprintf("%v:22", ipOrDom))
		userName = append(userName, user)
		password = append(password, pass)
		projectRoot = strings.ReplaceAll(projectRoot, "~", "")
		if user == "root" {
			projectRoot = filepath.Join("/root/", projectRoot)
		} else {
			projectRoot = filepath.Join("/home/"+user+"/", projectRoot)
		}

		projectPath = append(projectPath, projectRoot)
	}

	fmt.Printf("Project Directory From '~' (eg: ~/my/path): ")
	fmt.Scan(&projectPath)

	inventoryConfig := Inventory{
		UserName:       userName,
		PassWord:       password,
		NodeIpOrDomain: domainOrip,
		ProjectPath:    projectPath,
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
	err = os.WriteFile(ConfigFile, userConfigJson, 0660)

	if err != nil {
		fmt.Printf("Failed to write input data to %v\n", ConfigFile)
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
