package main

import (
	gintransport "example/auth-services/internal/auth/transport/gin"
	"net/http"

	"github.com/gin-gonic/gin"
)


var secret = []byte("secret")

func main() {
	r := gin.Default()

	// public route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// protected route
	auth := r.Group("/auth")
	{
		auth.POST("/login", login)
		auth.POST("/register", gintransport.)
	}



	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}

