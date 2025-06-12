package main

import (
	"github.com/Webglhost-QA-Backend/backend/config"
	"github.com/Webglhost-QA-Backend/backend/internal/routers"
	"github.com/Webglhost-QA-Backend/backend/pkg/cache_client"
	"github.com/Webglhost-QA-Backend/backend/pkg/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	router := gin.Default()
	cfg := config.LoadConfig()
	db, err := database.InitDB(&cfg.MYSQL)

	if err != nil {
		log.Fatalf("Failed to initialize database:%v", err)
	}
	defer database.CloseDB()

	mongo, err := database.InitMongo(&cfg.MONGO)
	if err != nil {
		log.Fatalf("Failed to initialize mongo:%v", err)
	}
	defer database.MongoClose()

	// 静态文件服务 - 指向你的dist文件夹
	router.Static("/assets", "./dist/assets")
	router.StaticFile("/favicon.ico", "./dist/favicon.ico")

	// 处理前端路由 - 所有未匹配的路径都返回index.html
	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if len(path) >= 3 && path[:3] == "/v/" {
			c.File("./dist/index.html") // 返回 Vue 的 index.html
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		}
	})

	redisClient := new(cache_client.Redis)
	redisClient.Cfg = cfg.REDIS
	redisClient.Connect()

	routers.SetupRouter(router, db, cfg, mongo, redisClient)

	go func() {
		if err := router.Run(cfg.SERVER.PORT); err != nil {
			log.Fatalf("Failed to start server:%v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

}
