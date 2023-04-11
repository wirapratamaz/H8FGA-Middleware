package controllers

import (
	"jwtgo/database"
	"jwtgo/helpers"
	"jwtgo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRole(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	Role := models.Role{}

	if contentType == appJson {
		c.ShouldBindJSON(&Role)
	} else {
		c.ShouldBind(&Role)
	}

	err := db.Debug().Create(&Role).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          Role.ID,
		"name":        Role.Name,
		"description": Role.Description,
	})
}
