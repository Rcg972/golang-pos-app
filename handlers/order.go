package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/taufik-s/project-go/models"
)

func init() {
	LoadOrdersFromFile()
}

var orders = []models.Order{}

func PostOrder(c *gin.Context) {
	if c.PostForm("item_name") == "" {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":    "Kedai Nusantara",
			"subtitle": "Selamat Datang di Kedai Nusantara",
			"menu":     menu,
			"error":    "Silahkan pilih item",
		})
		return
	} else if quantity, err := strconv.Atoi(c.PostForm("quantity")); err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":    "Kedai Nusantara",
			"subtitle": "Selamat Datang di Kedai Nusantara",
			"menu":     menu,
			"error":    "Quantity harus berupa angka",
		})
		return
	} else if quantity < 1 {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":    "Kedai Nusantara",
			"subtitle": "Selamat Datang di Kedai Nusantara",
			"menu":     menu,
			"error":    "Quantity harus lebih dari 0",
		})
		return
	} else {
		var harga int

		for _, item := range menu {
			if item.Name == c.PostForm("item_name") {
				harga = item.Price
				break
			}
		}

		pesanan := models.Order{
			ItemName: c.PostForm("item_name"),
			Quantity: quantity,
			Status:   "Berhasil",
			Total:    quantity * harga,
		}
		orders = append(orders, pesanan)
		saveOrdersToFile()
	}

	c.Redirect(http.StatusFound, "/orders")
}

func GetOrders(c *gin.Context) {
	c.HTML(http.StatusOK, "orders.html", gin.H{
		"orders": orders,
	})
}

func saveOrdersToFile() {
	data, err := json.MarshalIndent(orders, "", "  ")
	if err != nil {
		fmt.Println("Gagal Menyimpan Data", err)
		return
	}
	os.WriteFile("orders.json", data, 0644)
}

func LoadOrdersFromFile() {
	data, err := os.ReadFile("orders.json")
	if err != nil {
		fmt.Println("Gagal Membaca File", err)
		return
	}
	json.Unmarshal(data, &orders)
}
