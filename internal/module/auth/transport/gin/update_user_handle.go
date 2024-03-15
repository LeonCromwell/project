package gintransport

import (
	"example/auth-services/internal/module/auth/business"
	"example/auth-services/internal/module/auth/storage"
	"example/auth-services/internal/pkg/auth"
	"example/auth-services/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func UpdateUserHandle(db *gorm.DB) func(c *gin.Context) {
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

		storage := storage.NewStorage(db)
		userBusiness := business.UpdateUserBusiness(storage)

		var user *model.UpdateUserInput
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updatedUser, err := userBusiness.UpdateUserBusiness(*userID, *user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User updated",
			"user": updatedUser,
		})
	}
}