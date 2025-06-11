package controllers

import (
	"github.com/Webglhost-QA-Backend/backend/config"
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/internal/app/services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type GameController struct {
	GameService services.GameService
	MessageChan chan string
	config      config.Config
}

func NewGameController(gameService services.GameService, config config.Config) *GameController {
	return &GameController{
		GameService: gameService,
		MessageChan: make(chan string),
		config:      config,
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

func (gc *GameController) AddGame(c *gin.Context) {
	var GameInfo models.Game
	if err := c.ShouldBind(&GameInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}
	if err2 := gc.GameService.AddGame(GameInfo); err2 != nil {
		log.Printf("Faild to add game:%v", err2)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err2.Error(),
			"status":  false,
		})
		return
	}
	log.Println("Successful add game info")
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful add game info",
		"status":  true,
	})
}

func (gc *GameController) DeleteGame(c *gin.Context) {
	var GameInfo models.Game
	if err := c.ShouldBind(&GameInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}
	if err2 := gc.GameService.DeleteGame(GameInfo); err2 != nil {
		log.Printf("Faild to delete game:%v", err2)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err2.Error(),
			"status":  false,
		})
		return
	}
	log.Println("Successful delete game info")
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful delete game info",
		"status":  true,
	})
}

func (gc *GameController) WebSocket(c *gin.Context) {
	wsUpgrader := websocket.Upgrader{
		HandshakeTimeout: 10 * time.Second,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	go func() {
		for msg := range gc.MessageChan {
			ws.WriteMessage(websocket.TextMessage, []byte(msg))
		}
	}()

	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		switch messageType {
		case websocket.TextMessage:
			log.Printf("处理文本信息:%s\n", string(message))
			ws.WriteMessage(websocket.TextMessage, message)
		case websocket.CloseMessage:
			log.Println("关闭websocket")
			return
		default:
			log.Println("未知消息")
			return
		}
	}
}

func (gc *GameController) UpdateByFeishu(c *gin.Context) {
	go func() {
		gc.GameService.UpdateMongoByFeishu(&gc.MessageChan, gc.config.FEISHU)
	}()
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful get commond",
		"status":  true,
	})
}
