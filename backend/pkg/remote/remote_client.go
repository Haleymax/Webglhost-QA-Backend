package remote

import (
	"fmt"
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
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

	stdin, err := session.StdinPipe()
	if err != nil {
		return fmt.Errorf("create stdin pipe: %s", err)
	}
	defer stdin.Close()

	go func() {
		defer stdin.Close()

		if _, err := fmt.Fprintf(stdin, "C%04o %d %s\n", fileInfo.Mode().Perm(), fileInfo.Size(), filepath.Base(localPath)); err != nil {
			log.Printf("write to stdin: %s", err)
			return
		}

		if _, err := io.Copy(stdin, file); err != nil {
			log.Printf("copy file to stdin: %s", err)
			return
		}

		if _, err := fmt.Fprintf(stdin, "\x00"); err != nil {
			log.Printf("write end marker to stdin: %s", err)
			return
		}
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
	cmd := fmt.Sprintf("ls %s", remoteDir)
	output, err := r.RunCommadn(cmd)
	if err != nil {
		return nil, fmt.Errorf("run command: %s, output: %s", cmd, string(output))
	}
	lines := strings.Split(strings.TrimSpace(output), "\n")
	return lines, nil
}

func (r RemoteClient) DeleteAPKFiles(remoteDir string) error {
	if r.client == nil {
		return fmt.Errorf("client is not connected")
	}
	fileNames, err := r.GetRemoteFileNames(remoteDir)
	if err != nil {
		return err
	}
	for _, fileName := range fileNames {
		if strings.HasSuffix(fileName, ".apk") {
			cmd := fmt.Sprintf("rm  %s/%s", remoteDir, fileName)
			_, err := r.RunCommadn(cmd)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r RemoteClient) GetADBDevices(adb_path string) ([]string, error) {
	if r.client == nil {
		return nil, fmt.Errorf("client is not connected")
	}
	cmd := adb_path + "adb devices"
	output, err := r.RunCommadn(cmd)
	if err != nil {
		return nil, fmt.Errorf("run command: %s, output: %s", cmd, string(output))
	}
	lines := strings.Split(strings.TrimSpace(output), "\n")
	var devices []string

	for _, line := range lines {
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "List of devices") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) == 2 && fields[1] == "device" {
			devices = append(devices, fields[0])
		}
	}
	return devices, nil
}

func (r RemoteClient) getDeviceProperty(adb_path, serial, prop string) string {
	cmd := fmt.Sprintf("%sadb -s %s shell getprop %s", adb_path, serial, prop)
	result, err := r.RunCommadn(cmd)
	if err != nil {
		return "Unknown"
	}
	result = strings.TrimSpace(result)
	if len(result) < 2 {
		return "Unknown"
	}
	return result
}

func (r RemoteClient) getDeviceMarketName(adb_path, serial string) string {
	marketProps := []string{
		MarketNames.OPPO,
		MarketNames.HONOR,
		MarketNames.XIAOMI,
		MarketNames.IQOO,
		MarketNames.HUAWEI,
		MarketNames.ONEPLUS,
		MarketNames.REDMI,
	}

	for _, prop := range marketProps {
		result := r.getDeviceProperty(adb_path, serial, prop)
		if result != "Unknown" && len(result) > 1 {
			return result
		}
	}
	return "Market name not found"
}

func (r RemoteClient) GetPhoneInfo(adb_path, serial string) (models.Phone, error) {
	if r.client == nil {
		return models.Phone{}, fmt.Errorf("client is not connected")
	}

	manufacturer := r.getDeviceProperty(adb_path, serial, "ro.product.manufacturer")
	model := r.getDeviceProperty(adb_path, serial, "ro.product.model")
	androidVersion := r.getDeviceProperty(adb_path, serial, "ro.build.version.release")
	cpuAbi := r.getDeviceProperty(adb_path, serial, "ro.product.cpu.abi")
	marketName := r.getDeviceMarketName(adb_path, serial)

	phone := models.Phone{
		Manufacturer:   manufacturer,
		Model:          model,
		AndroidVersion: androidVersion,
		Cpuabi:         cpuAbi,
		MarketName:     marketName,
		Serial:         serial,
	}

	return phone, nil
}

func (r *RemoteClient) Close() {
	if err := r.client.Close(); err != nil {
		log.Printf("close remote client: %s", err)
	}
}
