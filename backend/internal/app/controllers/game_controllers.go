package controllers

import (
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
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

func (gc *GameController) FindAllWXGame(c *gin.Context) {
	if _, err := gc.GameService.FindAllWxGame(); err != nil {
		log.Printf("Faild to get gamelist:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Faild to get gamelist" + err.Error(),
			"games":   []map[string]interface{}{},
			"status":  false,
		})
		return
	}

	games, _ := gc.GameService.FindAllWxGame()
	log.Println("Successful get all game list")
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful get all game list",
		"games":   games,
		"status":  true,
	})
}

func (gc *GameController) FindAllRPK(c *gin.Context) {
	if _, err := gc.GameService.FindAllRPK(); err != nil {
		log.Printf("Faild to get gamelist:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Faild to get gamelist" + err.Error(),
			"games":   []map[string]interface{}{},
			"status":  false,
		})
		return
	}

	games, _ := gc.GameService.FindAllRPK()
	log.Println("Successful get all game list")
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful get all game list",
		"games":   games,
		"status":  true,
	})
}

func (gc *GameController) FindAllGameByType(c *gin.Context) {
	var GameType models.GameRequest
	if err := c.ShouldBind(&GameType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"games":   []map[string]interface{}{},
			"status":  false,
		})
		return
	}

	if _, err := gc.GameService.FindAllByType(GameType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"games":   []map[string]interface{}{},
			"status":  false,
		})
		return
	}
	games, _ := gc.GameService.FindAllByType(GameType)
	log.Println("Successful get all game list")
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful get all game list",
		"games":   games,
		"status":  true,
	})
}

func (gc *GameController) UpdateGameInfo(c *gin.Context) {
	var GameInfo models.Game
	if err := c.ShouldBind(&GameInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	if err2 := gc.GameService.UpdateGame(GameInfo); err2 != nil {
		log.Printf("Faild to update game:%v", err2)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err2.Error(),
			"status":  false,
		})
		return
	}

	log.Println("Successful update game info")
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful update game info",
		"status":  true,
	})
}
