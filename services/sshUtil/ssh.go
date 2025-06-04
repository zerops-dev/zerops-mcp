package sshUtil

import (
	"io"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func WriteFile(remoteHost string, srcFile io.Reader, dstPath string) error {
	config := &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, _ := ssh.Dial("tcp", remoteHost+":22", config)
	defer client.Close()

	// open an SFTP session over an existing ssh connection.
	sftp, err := sftp.NewClient(client)
	if err != nil {
		return err
	}
	defer sftp.Close()

	// Create the destination file
	dstFile, err := sftp.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	return nil
}

func ReadFile(remoteHost string, srcFile string, dts io.Writer) error {
	config := &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, _ := ssh.Dial("tcp", remoteHost+":22", config)
	defer client.Close()

	// open an SFTP session over an existing ssh connection.
	sftp, err := sftp.NewClient(client)
	if err != nil {
		return err
	}
	defer sftp.Close()

	// Create the destination file
	src, err := sftp.OpenFile(srcFile, os.O_RDONLY)
	if err != nil {
		return err
	}
	defer src.Close()

	if _, err := io.Copy(dts, src); err != nil {
		return err
	}

	return nil
}

func ReaddDirectory(remoteHost string, directory string) ([]string, error) {
	config := &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, _ := ssh.Dial("tcp", remoteHost+":22", config)
	defer client.Close()

	// open an SFTP session over an existing ssh connection.
	sftp, err := sftp.NewClient(client)
	if err != nil {
		return nil, err
	}
	defer sftp.Close()

	// Create the destination file
	files, err := sftp.ReadDir(directory)
	if err != nil {
		return nil, err
	}
	results := make([]string, len(files))
	for _, result := range files {
		results = append(results, result.Name())
	}

	return results, nil
}
