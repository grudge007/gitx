package pushx

import (
	"encoding/json"
	"fmt"
	"gitx/initx"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func PushFilesToRemote() {

	// unmarshall data
	var userConfig initx.Config
	var ignoreFile []string
	configFile, _ := os.ReadFile(".gitx/gitx.conf")
	err := json.Unmarshal(configFile, &userConfig)
	if err != nil {
		log.Fatalf("Failed to load configuration : %v\n", err)
	}
	count := userConfig.NumberOfNodes
	for i := 0; i < count; i++ {

		user := userConfig.InventoryStruct.UserName[i]
		password := userConfig.InventoryStruct.PassWord[i]
		aadr := userConfig.InventoryStruct.NodeIpOrDomain[i]
		projectPath := userConfig.InventoryStruct.ProjectPath[i]
		config := &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.Password(password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         5 * time.Second,
		}

		conn, err := ssh.Dial("tcp", aadr, config)
		if err != nil {
			log.Fatal("Failed to connect ssh")
			return
		}
		defer conn.Close()

		client, err := sftp.NewClient(conn)
		if err != nil {
			log.Fatal("Failed to create a new session")
			return
		}
		defer client.Close()

		data, _ := os.ReadFile(".gitx/ignore")
		lines := strings.Split(string(data), "\n")

		for _, iteralData := range lines {
			ignoreFile = append(ignoreFile, string(iteralData))

		}

		// fmt.Println(ignoreFile)
		projectRoot := "/home/iamgrudge/Configs/PE/portfolio/gitx"
		filepath.Walk(projectRoot, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			for _, iteralData := range ignoreFile {
				if info.Name() == iteralData {
					if info.IsDir() {
						return filepath.SkipDir
					}
					return nil
				}

			}
			relPath, _ := filepath.Rel(projectRoot, path)
			remotePath := filepath.Join(projectPath, relPath)

			if info.IsDir() {
				return client.MkdirAll(remotePath)
			}

			localFile, err := os.Open(path)
			if err != nil {
				log.Fatal("Failed to open ", info.Name())
				return nil
			}
			defer localFile.Close()

			remoteFile, err := client.Create(remotePath)
			if err != nil {
				log.Fatal("Failed to create ", path)
				return nil
			}
			defer remoteFile.Close()

			bytesCopied, err := io.Copy(remoteFile, localFile)
			if err != nil {
				log.Fatal("Failed to copy data ", info.Name())
				return nil
			}
			fmt.Printf("Successfully transferred %d bytes to remote server.\n", bytesCopied)
			return nil
		})
	}
}
