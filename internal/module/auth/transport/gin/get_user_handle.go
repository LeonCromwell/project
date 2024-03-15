package gintransport

import (
	"example/auth-services/internal/module/auth/business"
	"example/auth-services/internal/module/auth/storage"
	"example/auth-services/internal/pkg/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)



func GetUserHandler(db *gorm.DB) func(c *gin.Context) {
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

		store := storage.NewStorage(db)
		userBusiness := business.GetuserBusiness(store)

		user, err := userBusiness.GetUser(c.Request.Context(), *userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "User not found",
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}
}

