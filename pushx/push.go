package push

import (
	"encoding/json"
	"fmt"
	"gitx/initx"
	"io"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func PushFilesToRemote() {
	var userConfig initx.Config

	confData, err := os.ReadFile(initx.ConfigFile)
	if err != nil {
		fmt.Print("Error Opening Config File\n")
		return
	}
	err = json.Unmarshal(confData, &userConfig)
	if err != nil {
		fmt.Print("Error Opening Config File\n")
		return
	}

	for i := 0; i < userConfig.NumberOfNodes; i++ {
		ipOrDom := userConfig.InventoryStruct.NodeIpOrDomain[i]
		user := userConfig.InventoryStruct.UserName[i]
		pass := userConfig.InventoryStruct.PassWord[i]

		sshConfig := &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.Password(pass),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
		conn, err := ssh.Dial("tcp", ipOrDom, sshConfig)
		if err != nil {
			fmt.Print("Err: ", err)
			return
		}
		defer conn.Close()

		sftpClient, err := sftp.NewClient(conn)
		if err != nil {
			fmt.Print("Err: ", err)
			return
		}
		defer sftpClient.Close()
		file := "/tmp/test-go.txt"
		localFile, err := os.Open(file)
		if err != nil {
			fmt.Print("Err: ", err)
			return
		}
		defer localFile.Close()

		remoteFile, err := sftpClient.Create(file)
		if err != nil {
			fmt.Print("Err: ", err)
			return
		}

		bytes, err := io.Copy(remoteFile, localFile)
		if err != nil {
			fmt.Print("Err: ", err)
			return
		}
		defer remoteFile.Close()
		fmt.Printf("%d bytes copied successfully!\n", bytes)
	}

}
