package controllers

import (
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/internal/app/services"
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

	err := c.ShouldBindJSON(&node)
	log.Panicln(node.Host)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false,
		})
	}

	if err := nc.deviceService.AddNode(&node); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  true,
	})
}
