package services

import "github.com/Webglhost-QA-Backend/backend/config"

type RemoteService interface {
	UpLoad(filePath string, host string) error
}

type RemoteServiceImpl struct {
	RemoteDir string
}

func NewRemoteService(cfg *config.RemoteConfig) RemoteService {
	return &RemoteServiceImpl{RemoteDir: cfg.REMOTEDIR}
}

func (r RemoteServiceImpl) UpLoad(filePath string) error {
	return nil
}
