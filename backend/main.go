package main

import (
	"github.com/Webglhost-QA-Backend/backend/config"
	"github.com/Webglhost-QA-Backend/backend/internal/routers"
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

	routers.SetupRouter(router, db)

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
