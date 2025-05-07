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
			"message": "Failed to connect to node:" + err.Error(),
			"status":  false,
		})
		return
	}
	defer remoteClient.Close()

	if err := nc.deviceService.AddNode(&node); err != nil {
		log.Printf("Failed to add node %s: %v", node.Host, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to register node",
			"status":  false,
		})
		return
	}

	log.Printf("Successfully added node %s", node.Host)
	c.JSON(http.StatusOK, gin.H{
		"message": "Node added successfully",
		"status":  true,
	})
}

func (nc *NodeController) UpdataNode(c *gin.Context) {
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
			"message": "Failed to connect to node:" + err.Error(),
			"status":  false,
		})
		return
	}

	if err := nc.deviceService.UpdateNode(node); err != nil {
		log.Printf("Failed to update node %s: %v", node.Host, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to register node:" + err.Error(),
			"status":  false,
		})
		return
	}
	log.Printf("Successfully updated node %s", node.Host)
	c.JSON(http.StatusOK, gin.H{
		"message": "Node updated successfully",
		"status":  true,
	})
}

func (nc *NodeController) RemoveNode(c *gin.Context) {
	host := c.Query("host")
	if host == "" {
		log.Printf("Failed to get host from request")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"status":  false,
		})
		return
	}

	if err := nc.deviceService.DeleteNode(host); err != nil {
		log.Printf("Failed to remove node %s: %v", host, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to remove node:" + err.Error(),
			"status":  false,
		})
		return
	}
	log.Printf("Successfully removed node %s", host)
	c.JSON(http.StatusOK, gin.H{
		"message": "Node removed successfully",
		"status":  true,
	})
}

func (nc *NodeController) GetNodes(c *gin.Context) {
	if _, err := nc.deviceService.FindAllNodes(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get nodes" + err.Error(),
			"nodes":   []map[string]string{},
			"status":  false,
		})
		return
	}
	nodes, _ := nc.deviceService.FindAllNodes()
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully get nodes",
		"nodes":   nodes,
		"status":  true,
	})
}
