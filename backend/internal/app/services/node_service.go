package services

import (
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/internal/app/repositories"
)

type HostMap map[string]string

type NodeService interface {
	AddNode(node *models.Node) error
	DeleteNode(id string) error
	UpdateNode(node models.Node) error
	FindNode(host string) (*models.Node, error)
	FindAllNodes() ([]HostMap, error)
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

func (s *NodeServiceImpl) FindAllNodes() ([]HostMap, error) {
	nodes, err := s.deviceRepo.FindAllNodes()
	if err != nil {
		return nil, err
	}
	hostMaps := make([]HostMap, len(nodes))
	for i, node := range nodes {
		hostMaps[i] = HostMap{
			"host": node.Host,
			"name": node.Name,
		}
	}
	return hostMaps, nil
}
