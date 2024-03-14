package gintransport

import (
	"example/auth-services/internal/module/auth/business"
	"example/auth-services/internal/module/auth/storage"
	"example/auth-services/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func VertifyHandle(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var vertifyInput model.VertifyInput

		if err := c.ShouldBindJSON(&vertifyInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}

		store := storage.NewStorage(db)
		business := business.VertifyBusiness(store)

		if err := business.VertifyBusiness(c.Request.Context(), &vertifyInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Vertify success",
		})
	}
}
