package main

import (
	"fmt"
	"os"

	sftp "github.com/mmaous/sftp-uploader/utils"
)

func main() {

	localFolderPath := os.Getenv("LOCAL_FOLDER_PATH")
	remoteFolderPath := os.Getenv("REMOTE_FOLDER_PATH")
	username := os.Getenv("VM_USERNAME")
	password := os.Getenv("VM_PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "22"
	}
	
	client, err := sftp.SetupSSHClient(host, port, username, password)
	if err != nil {
		fmt.Printf("Failed to set up SSH client: %v\n", err)
		return
	}
	defer client.Close()

	err = sftp.ChangeDirectory(client, remoteFolderPath)
	if err != nil {
		fmt.Printf("Failed to change directory to %s on the remote server: %v\n", remoteFolderPath, err)
		return
	}

	transferredFilesCount, err := sftp.TransferFiles(client, localFolderPath, remoteFolderPath)
	if err != nil {
		fmt.Printf("Failed to transfer files: %v\n", err)
		return
	}

	fmt.Printf("Transferred %d files.\n", transferredFilesCount)
	fmt.Println("All files uploaded successfully.")
}

