package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/taufik-s/project-go/handlers"
)

func main() {

	router := gin.Default()
	router.Static("/static", "./assets")

	router.LoadHTMLFiles("templates/index.html", "templates/order.html", "templates/orders.html")

	router.GET("/", handlers.GetIndex)

	router.POST("/order", handlers.PostOrder)

	router.GET("/orders", handlers.GetOrders)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
