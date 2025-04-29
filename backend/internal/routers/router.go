package routers

import (
	"github.com/Webglhost-QA-Backend/backend/internal/app/controllers"
	"github.com/Webglhost-QA-Backend/backend/internal/app/repositories"
	"github.com/Webglhost-QA-Backend/backend/internal/app/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(router *gin.Engine, db *gorm.DB) {
	nodeRepo := repositories.NewNodeRepository(db)
	nodeService := services.NewNodeService(nodeRepo)
	nodeController := controllers.NewNodeController(nodeService)
	initControler := controllers.NewInitController(db)

	api := router.Group("/api/v1")
	init := api.Group("/init")
	{
		init.GET("/", initControler.Init)
	}
	nodes := api.Group("/nodes")
	{
		nodes.POST("/add", nodeController.AddNode)
	}
}
