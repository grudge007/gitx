package pushz

import (
	"fmt"
	"gitx/initz"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
)

var gitzConfig initz.Inventory

func PushFilesToRemote() {
	var wg sync.WaitGroup
	gitzConfigFile := "/home/iamgrudge/Configs/PE/portfolio/gitx/.gitz/gitz.yaml"
	nodeDetails := loadGitzConf(gitzConfigFile)
	ignorFiles := loadGitzIgnore()
	filesToSend := getFilesList(ignorFiles)
	pvtKey := getSshPubKey(nodeDetails.PrivateKeyPath)

	for i := 0; i < len(nodeDetails.Nodes); i++ {
		connection := getSshConnection(pvtKey, nodeDetails.Nodes[i].User, nodeDetails.Nodes[i].IP)

		wg.Add(1)
		go func(node initz.Node) {
			defer wg.Done()
			defer connection.Close()
			PushFiles(nodeDetails.ProjectRoot, connection, filesToSend, node.Path)
		}(nodeDetails.Nodes[i])

	}
	wg.Wait()

}

func loadGitzConf(gitzConfigFile string) initz.Inventory {
	data, err := os.ReadFile(gitzConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &gitzConfig)
	return gitzConfig
}

func loadGitzIgnore() []string {
	var ignores []string
	gitzIgnore := "/home/iamgrudge/Configs/PE/portfolio/gitx/.gitz/ignores"
	data, _ := os.ReadFile(gitzIgnore)
	lines := strings.Split(string(data), "\n")
	for _, file := range lines {
		ignores = append(ignores, string(file))
	}
	return ignores
}

func getFilesList(ignoreFiles []string) []string {
	var filesToSend []string
	root, _ := os.Getwd()
	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		for _, file := range ignoreFiles {
			if file == info.Name() {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}
		if info.IsDir() {
			return nil
		}
		filesToSend = append(filesToSend, path)
		return nil

	})
	if err != nil {
		fmt.Println(err)
	}
	return filesToSend
}

func getSshPubKey(pvtKeyPath string) ssh.Signer {
	pvtKey, err := os.ReadFile(pvtKeyPath)
	if err != nil {
		fmt.Println(err)
	}
	signer, err := ssh.ParsePrivateKey(pvtKey)
	if err != nil {
		fmt.Println(err)
	}
	return signer
}

func getSshConnection(pvtKey ssh.Signer, user string, node string) *ssh.Client {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(pvtKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	conn, err := ssh.Dial("tcp", node+":22", config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return conn
}

func PushFiles(projectRoot string, conn *ssh.Client, files []string, projectPath string) {
	createdDirs := make(map[string]bool)
	client, err := sftp.NewClient(conn)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()
	for _, path := range files {
		relPath, _ := filepath.Rel(projectRoot, path)
		remotePath := filepath.Join(projectPath, relPath)
		remoteDir := filepath.Dir(remotePath)
		if !createdDirs[remoteDir] {
			fmt.Printf("create dir %v\n", remoteDir)
			err := client.MkdirAll(remoteDir)
			if err != nil {
				fmt.Printf("Warning: could not create dir %v: %v\n", remoteDir, err)
			}
			createdDirs[remoteDir] = true

		}
		localFile, err := os.Open(path)
		if err != nil {
			fmt.Errorf("failed to read %v, skipping\n", path)
			continue
		}

		remoteFile, err := client.Create(remotePath)
		if err != nil {
			fmt.Errorf("failed to create %v, skipping\n", remotePath)
			continue
		}

		bytesCopied, err := io.Copy(remoteFile, localFile)
		if err != nil {
			fmt.Errorf("Failed to copy %v, skipping\n", path)
		} else {
			fmt.Printf("Successfully Copied %d bytes: %s\n", bytesCopied, relPath)

		}
		localFile.Close()
		remoteFile.Close()
	}

}
