package runx

import (
	"encoding/json"
	"fmt"
	"gitx/initx"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func RunCommand(remoteCommand string) {
	var userConfig initx.Config
	configFile, _ := os.ReadFile(".gitx/gitx.conf")
	err := json.Unmarshal(configFile, &userConfig)
	if err != nil {
		log.Fatal("Error Loading Configuration")
	}
	count := userConfig.NumberOfNodes
	for i := 0; i < count; i++ {
		user := userConfig.InventoryStruct.UserName[i]
		passwd := userConfig.InventoryStruct.PassWord[i]
		addr := userConfig.InventoryStruct.NodeIpOrDomain[i]

		config := &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.Password(passwd),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         5 * time.Second,
		}

		conn, err := ssh.Dial("tcp", addr, config)
		if err != nil {
			fmt.Errorf(err.Error())
			return
		}

		defer conn.Close()

		session, err := conn.NewSession()
		if err != nil {
			fmt.Errorf(err.Error())
			return
		}
		defer session.Close()
		output, err := session.CombinedOutput(remoteCommand)
		if err != nil {
			fmt.Errorf(err.Error())
		}
		fmt.Print(string(output))
	}
}
