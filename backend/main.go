package main

import (
	"github.com/Webglhost-QA-Backend/backend/config"
	"github.com/Webglhost-QA-Backend/backend/internal/routers"
	"github.com/Webglhost-QA-Backend/backend/pkg/cache_client"
	"github.com/Webglhost-QA-Backend/backend/pkg/database"
	"github.com/gin-gonic/gin"
	"log"
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
