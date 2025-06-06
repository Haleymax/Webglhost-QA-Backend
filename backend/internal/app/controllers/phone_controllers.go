package controllers

import (
	"github.com/Webglhost-QA-Backend/backend/internal/app/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PhoneController struct {
	PhoneService services.PhoneService
}

func NewPhoneController(phoneService services.PhoneService) PhoneController {
	return PhoneController{
		PhoneService: phoneService,
	}
}

func (pc *PhoneController) AddPhone(c *gin.Context) {
	
}

func (pc *PhoneController) FindAllPhone(c *gin.Context) {
	if _, err := pc.PhoneService.FindAllPhone(); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"phones":  []map[string]interface{}{},
			"status":  false,
		})
		return
	}

	phones, _ := pc.PhoneService.FindAllPhone()
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully get all phones",
		"phones":  phones,
		"status":  true,
	})
}
