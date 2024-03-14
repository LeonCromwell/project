package gintransport

import (
	"example/auth-services/internal/pkg/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)


func VerifyTokenHandle() func(c *gin.Context){
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Access-Token")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Token not found",
			})
			c.Abort()
			return
		}

		userID, err := auth.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Authorized",
			"user": userID,
		})

	}

}