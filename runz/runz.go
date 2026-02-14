package runz

import (
	"fmt"
	"gitx/initz"
	"os"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

type RunManager struct {
	Config initz.Inventory
	// RemoteCommand string
	Auth ssh.Signer
}

func RunCommand(loadedConfig *initz.Inventory, remoteCommand string) {
	var wg sync.WaitGroup

	runCommand := CommandExec(*loadedConfig)
	for i := 0; i < len(loadedConfig.Nodes); i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			fmt.Println(runCommand.cmdRunner(remoteCommand, index))

		}(i)

	}
	wg.Wait()

}

func CommandExec(inventory initz.Inventory) *RunManager {
	RunManager := &RunManager{
		Config: inventory,
	}
	// RemoteCommand =  getRemoteCommand(),
	RunManager.Auth = RunManager.getSshSigner()
	return RunManager
}

func (inventory *RunManager) getSshConnection(index int) *ssh.Client {
	config := &ssh.ClientConfig{
		User: inventory.Config.Nodes[index].User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(inventory.Auth),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}
	conn, err := ssh.Dial("tcp", inventory.Config.Nodes[index].IP+":22", config)
	if err != nil {
		fmt.Printf("[ERROR] SSH connection failed to %s: %v\n",
			inventory.Config.Nodes[index].IP, err)
		return nil
	}
	return conn
}

func (inventory *RunManager) getSshSigner() ssh.Signer {
	pvtKey, err := os.ReadFile(inventory.Config.PrivateKeyPath)
	if err != nil {
		fmt.Printf("[ERROR] Unable to read private key at %s: %v\n",
			inventory.Config.PrivateKeyPath, err)
		return nil
	}
	signer, err := ssh.ParsePrivateKey(pvtKey)
	if err != nil {
		fmt.Printf("[ERROR] Failed to parse private key at %s: %v\n",
			inventory.Config.PrivateKeyPath, err)
	}
	return signer
}

func (inventory *RunManager) cmdRunner(remoteCommand string, index int) string {
	connection := inventory.getSshConnection(index)
	if connection == nil {
		return fmt.Sprintf("[%s] SSH connection could not be established",
			inventory.Config.Nodes[index].IP)
	}
	defer connection.Close()

	session, err := connection.NewSession()
	if err != nil {
		connection.Close()
		return fmt.Sprintf("[ERROR] Failed to create SSH session on %s: %v",
			inventory.Config.Nodes[index].IP, err)
	}
	defer session.Close()

	output, err := session.CombinedOutput(remoteCommand)
	if err != nil {
		fmt.Printf("[ERROR] Command '%s' failed on %s: %v\n",
			remoteCommand,
			inventory.Config.Nodes[index].IP,
			err)
	}
	return string(output)
}
