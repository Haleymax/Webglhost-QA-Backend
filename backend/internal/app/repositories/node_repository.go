package repositories

import (
	"fmt"
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/pkg/database"
	"gorm.io/gorm"
	"log"
)

type NodeRepository interface {
	Create(node *models.Node) error
	FindByHost(host string) (*models.Node, error)
	Update(node models.Node) error
	DeleteByHost(host string) error
}

type NodeRepositoryImpl struct {
	db *gorm.DB
}

func NewNodeRepository(db *gorm.DB) *NodeRepositoryImpl {
	return &NodeRepositoryImpl{db: database.DB}
}

func (r *NodeRepositoryImpl) Create(node *models.Node) error {
	var existNode models.Node
	err := r.db.Where("host = ?", node.Host).First(&existNode).Error

	if err == gorm.ErrRecordNotFound {
		if err = r.db.Create(node).Error; err != nil {
			log.Fatal(err)
		}
		log.Println("node created")
		return nil
	} else if err != nil {
		log.Fatal(err)
		return err
	} else {
		log.Println("node already exists")
		return fmt.Errorf("node already exists")
	}
}

func (r *NodeRepositoryImpl) FindByHost(host string) (*models.Node, error) {
	var node models.Node
	err := r.db.Where("host = ?", host).First(&node).Error
	return &node, err
}

func (r *NodeRepositoryImpl) Update(node models.Node) error {
	var existNode models.Node
	err := r.db.Where("host = ?", node.Host).First(&existNode).Error
	if err == gorm.ErrRecordNotFound {
		return fmt.Errorf("node doesn't exist")
	} else if err != nil {
		return err
	} else {
		err = r.db.Model(&existNode).Updates(&node).Error
		return err
	}
}

func (r *NodeRepositoryImpl) DeleteByHost(host string) error {
	var node models.Node
	err := r.db.Where("host = ?", host).First(&node).Error
	if err == gorm.ErrRecordNotFound {
		return fmt.Errorf("node doesn't exist")
	} else if err != nil {
		return err
	} else {
		err = r.db.Delete(&node).Error
		return err
	}
}
