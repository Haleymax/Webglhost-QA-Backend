package services

import (
	"github.com/Webglhost-QA-Backend/backend/config"
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/pkg/remote"
	"log"
)

type RemoteService interface {
	UpLoad(filePath string, node models.Node) error
	GetPhone(node models.Node) ([]string, error)
}

type RemoteServiceImpl struct {
	RemoteDir string
}

func NewRemoteService(cfg *config.RemoteConfig) RemoteService {
	return &RemoteServiceImpl{RemoteDir: cfg.REMOTEDIR}
}

func (r RemoteServiceImpl) UpLoad(filePath string, node models.Node) error {
	remote_client := remote.NewRemoteClient(node.Host, 22, node.User, node.Password)
	err := remote_client.Connect()
	if err != nil {
		return err
	}
	defer remote_client.Close()

	// 清除远程目录中的apk文件
	if err = remote_client.DeleteAPKFiles(r.RemoteDir); err != nil {
		log.Printf("delete apk files error: %v", err)
		return err
	}

	// 通过scp将上传的文件传输给远程服务器
	if err = remote_client.SCPUPload(filePath, r.RemoteDir); err != nil {
		log.Printf("scp upload error: %v", err)
	}
	return nil

}

func (r RemoteServiceImpl) GetPhone(node models.Node) ([]string, error) {
	remote_client := remote.NewRemoteClient(node.Host, 22, node.User, node.Password)
	err := remote_client.Connect()
	if err != nil {
		return nil, err
	}
	defer remote_client.Close()

	// 获取远程节点中的手机
	var phones []string
	phones, err = remote_client.GetADBDevices()
	if err != nil {
		return nil, err
	}
	return phones, nil
}
