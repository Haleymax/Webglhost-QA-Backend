package controllers

import (
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/internal/app/services"
	"github.com/Webglhost-QA-Backend/backend/pkg/remote"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type NodeController struct {
	deviceService services.NodeService
}

func NewNodeController(deviceService services.NodeService) *NodeController {
	return &NodeController{deviceService: deviceService}
}

func (nc *NodeController) AddNode(c *gin.Context) {
	var node models.Node

	if err := c.ShouldBindJSON(&node); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"status":  false,
		})
		return
	}

	if node.Host == "" || node.User == "" || node.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing required fields",
			"status":  false,
		})
		return
	}

	remoteClient := remote.NewRemoteClient(node.Host, 22, node.User, node.Password)
	if err := remoteClient.Connect(); err != nil {
		log.Printf("Failed to connect to %s: %v", node.Host, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to connect to node",
			"status":  false,
			"details": err.Error(),
		})
		return
	}
	defer remoteClient.Close()

	if err := nc.deviceService.AddNode(&node); err != nil {
		log.Printf("Failed to add node %s: %v", node.Host, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to register node",
			"status":  false,
			"details": err.Error(),
		})
		return
	}

	log.Printf("Successfully added node %s", node.Host)
	c.JSON(http.StatusOK, gin.H{
		"message": "Node added successfully",
		"status":  true,
	})
}
