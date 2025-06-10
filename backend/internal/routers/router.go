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
	gameRepo := repositories.NewGamesRepository(mongo)

	nodeService := services.NewNodeService(nodeRepo)
	remoteService := services.NewRemoteService(&cfg.REMOTE)
	watcherService := services.NewWatcherService(watcherRepo, cache)
	phoneService := services.NewPhoneService(phoneRepo)
	gameService := services.NewGameService(gameRepo)

	nodeController := controllers.NewNodeController(nodeService, remoteService)
	watcherController := controllers.NewWatcherController(watcherService)
	phoneController := controllers.NewPhoneController(phoneService)
	gameController := controllers.NewGameController(gameService)

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
		phone.POST("/add", phoneController.AddPhone)
		phone.GET("/find", phoneController.FindAllPhone)
		phone.DELETE("/remove", phoneController.DeletePhone)
		phone.PUT("/update", phoneController.UpdatePhone)
	}

	game := api.Group("/game")
	{
		game.GET("/find_all", gameController.FindAllGames)
		game.GET("/find_wx", gameController.FindAllWXGame)
		game.GET("/find_rpk", gameController.FindAllRPK)
		game.POST("/find_by_type", gameController.FindAllGameByType)
		game.PUT("/update", gameController.UpdateGameInfo)
		game.POST("/add", gameController.AddGame)
		game.DELETE("/remove", gameController.DeleteGame)
		game.GET("/ws", gameController.WebSocket)
	}
}
