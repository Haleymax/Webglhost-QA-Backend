package controllers

import (
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/pkg/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InitController struct {
	db *gorm.DB
}

func NewInitController(db *gorm.DB) *InitController {
	return &InitController{db: database.DB}
}

func (ic *InitController) Init(c *gin.Context) {
	err := ic.db.AutoMigrate(&models.Node{})
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
