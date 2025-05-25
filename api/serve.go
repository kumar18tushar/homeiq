package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

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
			"message": "Hello, World from Gin on Vercel!",
		})
	})

	// Serve the request using the Gin router
	router.ServeHTTP(w, r)
}
