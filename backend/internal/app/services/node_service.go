package services

import (
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/internal/app/repositories"
)

type NodeService interface {
	AddNode(node *models.Node) error
	DeleteNode(id string) error
	UpdateNode(node models.Node) error
	FindNode(id string) (*models.Node, error)
}

type NodeServiceImpl struct {
	deviceRepo repositories.NodeRepository
}

func NewNodeService(deviceRepo repositories.NodeRepository) *NodeServiceImpl {
	return &NodeServiceImpl{deviceRepo: deviceRepo}
}

func (s *NodeServiceImpl) AddNode(node *models.Node) error {
	return s.deviceRepo.Create(node)
}

func (s *NodeServiceImpl) DeleteNode(host string) error {
	return s.deviceRepo.DeleteByHost(host)
}

func (s *NodeServiceImpl) UpdateNode(node models.Node) error {
	return s.deviceRepo.Updata(node)
}

func (s *NodeServiceImpl) FindNode(host string) (*models.Node, error) {
	return s.deviceRepo.FindByHost(host)
}
