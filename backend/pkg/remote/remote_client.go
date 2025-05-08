package remote

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type RemoteClient struct {
	host     string
	port     int
	user     string
	password string
	client   *ssh.Client
}

func NewRemoteClient(host string, port int, user string, password string) *RemoteClient {
	return &RemoteClient{host: host, port: port, user: user, password: password}
}

func (r *RemoteClient) Connect() error {
	config := &ssh.ClientConfig{
		User: r.user,
		Auth: []ssh.AuthMethod{
			ssh.Password(r.password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", r.host, r.port), config)
	if err != nil {
		return fmt.Errorf("dial: %s", err)
	}
	r.client = client
	return nil
}

func (r RemoteClient) SCPUPload(localPath string, remotePath string) error {
	if r.client == nil {
		return fmt.Errorf("client is not connected")
	}

	file, err := os.Open(localPath)
	if err != nil {
		return fmt.Errorf("open file: %s", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("get file info: %s", err)
	}

	session, err := r.client.NewSession()
	if err != nil {
		return fmt.Errorf("create session: %s", err)
	}
	defer session.Close()

	remoteDir := filepath.Dir(remotePath)

	if _, err := r.RunCommadn(fmt.Sprintf("mkdir -p %s", remoteDir)); err != nil {
		return fmt.Errorf("remote mkdir: %s", err)
	}

	go func() {
		w, _ := session.StdinPipe()
		defer w.Close()

		fmt.Fprintln(w, "C%04o %d %s\n", fileInfo.Mode().Perm(), fileInfo.Size(), filepath.Base(localPath))

		io.Copy(w, file)

		fmt.Fprintf(w, "\x00")
	}()

	if err := session.Run(fmt.Sprintf("scp -t %s", remotePath)); err != nil {
		return fmt.Errorf("scp: %s", err)
	}
	return nil
}

func (r RemoteClient) RunCommadn(cmd string) (string, error) {
	if r.client == nil {
		return "", fmt.Errorf("client is not connected")
	}
	session, err := r.client.NewSession()
	if err != nil {
		return "", fmt.Errorf("create session: %s", err)
	}
	defer session.Close()

	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", fmt.Errorf("run command: %s, output: %s", cmd, string(output))
	}
	return string(output), nil
}

func (r RemoteClient) GetRemoteFileNames(remoteDir string) ([]string, error) {
	if r.client == nil {
		return nil, fmt.Errorf("client is not connected")
	}
	cmd := fmt.Sprintf("ls -l %s", remoteDir)
	output, err := r.RunCommadn(cmd)
	if err != nil {
		return nil, fmt.Errorf("run command: %s, output: %s", cmd, string(output))
	}
	lines := strings.Split(strings.TrimSpace(output), "\n")
	return lines, nil
}

func (r *RemoteClient) Close() error {
	return r.client.Close()
}
