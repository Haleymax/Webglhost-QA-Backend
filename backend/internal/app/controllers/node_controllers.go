package controllers

import (
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/internal/app/services"
	"github.com/Webglhost-QA-Backend/backend/pkg/remote"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type NodeController struct {
	deviceService services.NodeService
	remoteService services.RemoteService
}

func NewNodeController(deviceService services.NodeService, remoteService services.RemoteService) *NodeController {
	return &NodeController{deviceService: deviceService, remoteService: remoteService}
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

func (nc *NodeController) Upload(c *gin.Context) {

	file, err := c.FormFile("file")
	host := c.PostForm("host")
	if err != nil || host == "" {
		log.Printf("Failed to get file: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"status":  false,
		})
		return
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".apk" {
		log.Printf("Failed to get file, not an apk file")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"status":  false,
		})
		return
	}

	current_path, err := os.Getwd()
	if err != nil {
		log.Printf("Failed to get current path: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get current path",
			"status":  false,
		})
		return
	}
	media_path := filepath.Join(current_path, "media")
	now := time.Now()
	formatted_time := now.Format("2006-01-02-15-04-05")
	file_name := formatted_time + "-" + file.Filename
	save_paht := filepath.Join(media_path, file_name)
	if err := c.SaveUploadedFile(file, save_paht); err != nil {
		log.Printf("Failed to save file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save file",
			"status":  false,
		})
		return
	}

	node, _ := nc.deviceService.FindNode(host)

	if node == nil {
		log.Printf("Failed to find node %s", host)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to find node",
			"status":  false,
		})
		return
	}

	err = nc.remoteService.UpLoad(save_paht, *node)
	if err != nil {
		log.Printf("Failed to upload file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to upload file",
			"status":  false,
		})
		return
	}
	err = os.Remove(save_paht)
	if err != nil {
		log.Printf("Failed to remove local file: %v, but successful load file", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to remove local file, but successful load file",
			"status":  false,
		})
		return
	}
	log.Printf("Successfully uploaded file %s", save_paht)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully uploaded file: " + file_name,
		"status":  true,
	})

}

func (nc *NodeController) GetADBDevices(c *gin.Context) {
	host := c.Query("host")
	if host == "" {
		log.Printf("Failed to get host from request")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"status":  false,
		})
		return
	}

	nodes, err := nc.deviceService.FindNode(host)
	if err != nil {
		log.Printf("Failed to get nodes: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get nodes",
			"status":  false,
		})
		return
	}
	var phones []string
	phones, err = nc.remoteService.GetPhone(*nodes)
	if err != nil {
		log.Printf("Failed to get phones: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get phones",
			"status":  false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully get phones",
		"nodes":   phones,
		"status":  true,
	})
}
