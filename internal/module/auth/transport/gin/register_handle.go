package gintransport

import (
	"example/auth-services/internal/module/auth/business"
	"example/auth-services/internal/module/auth/storage"
	"example/auth-services/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterHandle(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.UserInput

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}

		store := storage.NewStorage(db)

		business := business.RegisterBusiness(store)

		if err := business.RegisterBusiness(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Send code successfully",
		})

	}

}
