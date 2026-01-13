package controllers

import (
	"time"
	models "week2/Models"

	"github.com/gin-gonic/gin"
)

var artists = []models.Artist{
	{
		ID:       0,
		Name:     "Nidji",
		Email:    "nidji@gmail.com",
		Password: "password",
		Created:  time.Now().Format("2006-01-02 15:04:05"),
	},
}

func GetArtists(c *gin.Context) {
	c.JSON(200, artists)
}
