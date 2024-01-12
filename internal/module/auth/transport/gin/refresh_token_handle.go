package gintransport

import (
	"example/auth-services/internal/pkg/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)



func RefreshTokenHandle() func(c *gin.Context){
	return func(c *gin.Context){
		 refreshToken := c.Request.Header.Get("Refresh-token")
		 if refreshToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "refreshToken is required"})
			return
		 }

		newAccessToken, err := auth.RefreshToken(refreshToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "refreshToken is wrong"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": newAccessToken,
		})
	}
}