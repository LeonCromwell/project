package gintransport

import (
	"example/auth-services/internal/module/auth/business"
	"example/auth-services/internal/module/auth/storage"
	"example/auth-services/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func LoginHandle(db *gorm.DB) func(c *gin.Context){
	return func(c *gin.Context){
		var  user *model.UserLoginInput

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewStorage(db)
		business := business.LoginBusiness(store)

		var token,refreshToken, err = business.Login(c.Request.Context(), user)
		if err != nil {
			// if err.Error() == "User not active" {
			// 	c.JSON(http.StatusBadRequest, gin.H{"error": "User not active"})
			// 	return
			// }
			// c.JSON(http.StatusBadRequest, gin.H{"error": "Email or password is wrong"})
			// return
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}	

		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"refreshToken": refreshToken,
		})
		
	}
}