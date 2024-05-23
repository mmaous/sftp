package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func SetupSSHClient(host, port, username, password string) (*sftp.Client, error) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s:%s: %v", host, port, err)
	}

	client, err := sftp.NewClient(conn)
	if err != nil {
		return nil, fmt.Errorf("failed to create SFTP client: %v", err)
	}

	return client, nil
}

func ChangeDirectory(client *sftp.Client, remoteFolderPath string) error {
/*	err := client.Mkdir(remoteFolderPath)
	if err != nil {
		return fmt.Errorf("failed to create directory %s on the remote server: %v", remoteFolderPath, err)
	}
*/
	return nil
}

func TransferFiles(client *sftp.Client, localFolderPath, remoteFolderPath string) (int, error) {
	transferredFilesCount := 0

	err := filepath.Walk(localFolderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			localFilePath := path
			filename := filepath.Base(localFilePath)
			remoteFilePath := filepath.Join(remoteFolderPath, filename)

			localFile, err := os.Open(localFilePath)
			if err != nil {
				return err
			}
			defer localFile.Close()

			remoteFile, err := client.Create(remoteFilePath)
			if err != nil {
				return err
			}
			defer remoteFile.Close()

			_, err = io.Copy(remoteFile, localFile)
			if err != nil {
				return err
			}

			fmt.Printf("Uploaded %s to %s\n", filename, remoteFilePath)
			transferredFilesCount++
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return transferredFilesCount, nil
}
