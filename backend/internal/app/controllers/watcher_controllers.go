package controllers

import (
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/internal/app/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type WatcherController struct {
	WatcherService services.WatcherService
}

func NewWatcherController(watcherService services.WatcherService) *WatcherController {
	return &WatcherController{
		WatcherService: watcherService,
	}
}

func (wc *WatcherController) FindAllWatchers(c *gin.Context) {
	if _, err := wc.WatcherService.FindAllWatchers(); err != nil {
		log.Printf("fail to query data %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Faild to get watcher" + err.Error(),
			"watcher": []map[string]string{},
			"status":  false,
		})
		return
	}

	watchers, _ := wc.WatcherService.FindAllWatchers()

	log.Printf("successful query data %v", watchers)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully get watchers",
		"watcher": watchers,
		"status":  true,
	})
}

func (wc *WatcherController) AddWatcher(c *gin.Context) {
	var watcher models.Watcher
	if err := c.BindJSON(&watcher); err != nil {
		log.Printf("fail to bind data: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Faild to get watcher" + err.Error(),
			"status":  false,
		})
		return
	}

	if watcher.Name == "" || watcher.Resource == "" || watcher.Click == "" || watcher.Brand == nil || watcher.Tag == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "no empty parameters are allowed",
			"status":  false,
		})
		return
	}

	exist, err := wc.WatcherService.FindOneWatcher(watcher.Resource)
	if err != nil {
		log.Printf("fail to query data %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Faild to get watcher" + err.Error(),
			"status":  false,
		})
		return
	}
	if exist != nil {
		log.Printf("data %v, exists, cant add", exist.Resource)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "data" + exist.Resource + ", exists, cant add",
			"status":  false,
		})
		return
	}
	err = wc.WatcherService.AddWatcher(&watcher)
	if err != nil {
		log.Printf("fail to add watcher %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Faild to add watcher" + err.Error(),
			"status":  false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully add watcher",
		"status":  true,
	})

}

func (wc *WatcherController) UpdataWatcher(c *gin.Context) {
	var watcher models.Watcher
	if err := c.BindJSON(&watcher); err != nil {
		log.Printf("fail to bind data: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Faild to get watcher" + err.Error(),
			"status":  false,
		})
		return
	}

	if watcher.Name == "" || watcher.Resource == "" || watcher.Click == "" || watcher.Brand == nil || watcher.Tag == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "no empty parameters are allowed",
			"status":  false,
		})
		return
	}

	exist, err := wc.WatcherService.FindOneWatcher(watcher.Resource)
	if err != nil {
		log.Printf("fail to query data %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Faild to get watcher" + err.Error(),
			"status":  false,
		})
		return
	}
	if exist == nil {
		log.Printf("data %v, dont exists, cant updata", exist.Resource)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "data" + exist.Resource + ", dont exists, cant updata",
			"status":  false,
		})
		return
	}
	err = wc.WatcherService.UpdateWatcher(&watcher)
	if err != nil {
		log.Printf("fail to update watcher %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Faild to update watcher" + err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully update watcher",
		"status":  true,
	})
}
