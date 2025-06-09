package controllers

import (
	"github.com/Webglhost-QA-Backend/backend/internal/app/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type GameController struct {
	GameService services.GameService
}

func NewGameController(gameService services.GameService) *GameController {
	return &GameController{
		GameService: gameService,
	}
}

func (gc *GameController) FindAllGames(c *gin.Context) {
	if _, err := gc.GameService.FindAllGame(); err != nil {
		log.Printf("Faild to get gamelist:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Faild to get gamelist" + err.Error(),
			"games":   []map[string]interface{}{},
			"status":  false,
		})
		return
	}

	games, _ := gc.GameService.FindAllGame()
	log.Println("Successful get all game list")
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful get all game list",
		"games":   games,
		"status":  true,
	})
}
