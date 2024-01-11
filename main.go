package main

import (
	gintransport "example/auth-services/internal/module/auth/transport/gin"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var secret = []byte("secret")

func main() {
	err := godotenv.Load("app.env")
	if err != nil {
		fmt.Println(err.Error())
		return
	}



	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/auth_go_api?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	r := gin.Default()

	// protected route
	auth := r.Group("/auth")
	{
		auth.POST("/login")
		auth.POST("/register", gintransport.RegisterHandle(db))
		auth.POST("/vertify", gintransport.VertifyHandle(db))
		auth.POST("/vertify/sendcode", gintransport.SendCodeHandle(db))
	}

	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}
