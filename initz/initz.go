package initz

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Node struct {
	IP   string `yaml:"ip"`
	User string `yaml:"user"`
	Path string `yaml:"path"`
}

type Inventory struct {
	ProjectName    string `yaml:"project_name"`
	ProjectRoot    string `yaml:"project_root"`
	DefaultUser    string `yaml:"default_user"`
	DefaultPath    string `yaml:"default_path"`
	PrivateKeyPath string `yaml:"private_key_path"`
	Nodes          []Node `yaml:"nodes"`
}

func InitGitz(override string) {
	gitzConfFile, workingDir := workingDirAndConfigFile()
	fmt.Println(gitzConfFile)
	isExist := checkIsGitzExist(gitzConfFile, override)
	if isExist {
		return
	}

	inventoryFile := Inventory{
		ProjectName:    "MyProject",
		ProjectRoot:    workingDir,
		DefaultUser:    "root",
		DefaultPath:    "/opt/MyProject",
		PrivateKeyPath: getSshPvtKeyPath(),
		Nodes: []Node{
			{
				IP:   "192.168.100.10",
				User: "root",
				Path: "/root/mypath",
			},
			{
				IP:   "192.168.100.11",
				User: "user",
				Path: "/opt/mypath",
			},
		},
	}
	inventoryYaml := convertToYaml(inventoryFile)
	saveToGitzConf(inventoryYaml, gitzConfFile)

}

func workingDirAndConfigFile() (string, string) {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(".gitz", 0766)
	gitzconfigFile := filepath.Join(workingDir, ".gitz/gitz.yaml")
	return gitzconfigFile, workingDir
}

func checkIsGitzExist(gitzConfFile string, override string) bool {
	if override == "--force" || override == "-f" {
		return false
	}
	_, err := os.Stat(gitzConfFile)
	if err != nil {
		return false
	} else {
		fmt.Println("An Existing COnfiguration Found")
		fmt.Println("run with --force to override")
		return true
	}

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

func getSshPvtKeyPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "YOUR_SSH_IDRSA_FILE"
	}
	privateKeyPath := filepath.Join(homeDir, ".ssh/id_rsa")
	return privateKeyPath
}
