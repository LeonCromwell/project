package gintransport

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterHandle(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var input RegisterInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		if err := Register(db, input); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
		})
	
}}