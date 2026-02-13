package runx

func RunCommand(remoteCommand string) {
	// var userConfig initx.Config
	// configFile, err := os.ReadFile(".gitx/gitx.conf")
	// if err != nil {
	// 	fmt.Println("Error Loading Configuration")
	// }
	// err = json.Unmarshal(configFile, &userConfig)
	// if err != nil {
	// 	fmt.Println("Error Loading Configuration")
	// }
	// count := userConfig.NumberOfNodes
	// for i := 0; i < count; i++ {
	// 	user := userConfig.InventoryStruct.UserName[i]
	// 	passwd := userConfig.InventoryStruct.PassWord[i]
	// 	addr := userConfig.InventoryStruct.NodeIpOrDomain[i]

	// 	config := &ssh.ClientConfig{
	// 		User: user,
	// 		Auth: []ssh.AuthMethod{
	// 			ssh.Password(passwd),
	// 		},
	// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	// 		Timeout:         5 * time.Second,
	// 	}

	// 	conn, err := ssh.Dial("tcp", addr, config)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		// return
	// 		continue
	// 	}

	// 	// defer conn.Close()

	// 	session, err := conn.NewSession()
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		conn.Close()
	// 		// return
	// 		continue
	// 	}
	// 	// defer session.Close()
	// 	output, err := session.CombinedOutput(remoteCommand)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	fmt.Print(string(output))
	// 	session.Close()
	// 	conn.Close()

	// }
}
