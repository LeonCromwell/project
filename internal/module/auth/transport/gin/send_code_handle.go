package gintransport

import (
	"example/auth-services/internal/module/auth/business"
	"example/auth-services/internal/module/auth/storage"
	"example/auth-services/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SendCodeHandle(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.VertifyInput

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}

		store := storage.NewStorage(db)
		business := business.SendCodeBusiness(store)

		if err := business.SendCode(c.Request.Context(), &data); err != nil {
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
