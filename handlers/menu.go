package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taufik-s/project-go/models"
)

var menu = []models.MenuItem{
	{Name: "Espresso", Price: 30000},
	{Name: "Latte", Price: 35000},
	{Name: "Cappuccino", Price: 35000},
	{Name: "Americano", Price: 32000},
	{Name: "Flat White", Price: 33000},
	{Name: "Piccolo Latte", Price: 31000},
}

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":    "Kedai Nusantara",
		"subtitle": "Selamat Datang di Kedai Nusantara",
		"menu":     menu,
	})
}
