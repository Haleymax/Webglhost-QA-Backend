package controllers

import (
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
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
	var phone models.Phone
	if err := c.ShouldBindJSON(&phone); err != nil {
		log.Printf("Faild to bind Json:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"messsage": err.Error(),
			"status":   false,
		})
		return
	}

	ret, message := pc.PhoneService.AddPhone(phone)
	if !ret {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": message,
			"status":  false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"status":  true,
	})
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

func (pc *PhoneController) DeletePhone(c *gin.Context) {
	serial := c.Query("serial")
	if err := pc.PhoneService.DeletePhone(serial); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully delete phones",
		"status":  true,
	})
}

func (pc *PhoneController) UpdatePhone(c *gin.Context) {
	var phone models.Phone

	if err := c.ShouldBindJSON(&phone); err != nil {
		log.Printf("Faild to bind Json: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	if err := pc.PhoneService.UpdatePhone(&phone); err != nil {
		log.Printf("Update phone happen error %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully update phones",
		"status":  true,
	})
}
