package remote

import (
	"fmt"
	"golang.org/x/crypto/ssh"
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

func (r *RemoteClient) Close() error {
	return r.client.Close()
}
