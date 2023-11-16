package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./static")

	router.LoadHTMLGlob("templates/*")

	router.GET("/menu", func(c *gin.Context) {
		c.HTML(http.StatusOK, "menu.html", nil)
	})

	router.Run(":8080")
}
