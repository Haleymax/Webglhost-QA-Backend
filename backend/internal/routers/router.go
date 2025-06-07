package routers

import (
	"github.com/Webglhost-QA-Backend/backend/config"
	"github.com/Webglhost-QA-Backend/backend/internal/app/controllers"
	"github.com/Webglhost-QA-Backend/backend/internal/app/repositories"
	"github.com/Webglhost-QA-Backend/backend/internal/app/services"
	"github.com/Webglhost-QA-Backend/backend/pkg/cache_client"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func SetupRouter(router *gin.Engine, db *gorm.DB, cfg *config.Config, mongo *mongo.Client, cache *cache_client.Redis) {
	nodeRepo := repositories.NewNodeRepository(db)
	watcherRepo := repositories.NewWatcherRepository(mongo)
	phoneRepo := repositories.NewPhoneRepository(mongo)

	nodeService := services.NewNodeService(nodeRepo)
	remoteService := services.NewRemoteService(&cfg.REMOTE)
	watcherService := services.NewWatcherService(watcherRepo, cache)
	phoneService := services.NewPhoneService(phoneRepo)

	nodeController := controllers.NewNodeController(nodeService, remoteService)
	watcherController := controllers.NewWatcherController(watcherService)
	phoneController := controllers.NewPhoneController(phoneService)

	initControler := controllers.NewInitController(db)

	api := router.Group("/api/v1")
	init := api.Group("/init")
	{
		init.GET("/", initControler.Init)
	}
	nodes := api.Group("/nodes")
	{
		nodes.POST("/add", nodeController.AddNode)
		nodes.PUT("/updata", nodeController.UpdataNode)
		nodes.DELETE("/remove", nodeController.RemoveNode)
		nodes.GET("/get", nodeController.GetNodes)
		nodes.POST("/upload", nodeController.Upload)
		nodes.GET("/get_phone", nodeController.GetADBDevices)
		nodes.POST("/phone_info", nodeController.GetPhoneInfo)
	}
	watchers := api.Group("/watchers")
	{
		watchers.GET("/find", watcherController.FindAllWatchers)
		watchers.POST("/add", watcherController.AddWatcher)
		watchers.PUT("/update", watcherController.UpdataWatcher)
		watchers.DELETE("/delete", watcherController.DeleteWatcher)
		watchers.POST("/refresh", watcherController.RefreshCache)
	}

	phone := api.Group("/phone")
	{
		phone.GET("/find", phoneController.FindAllPhone)
		phone.DELETE("/remove", phoneController.DeletePhone)
	}
}
