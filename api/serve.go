package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Gin on Vercel!",
		})
	})

	router.ServeHTTP(w, r)
}
