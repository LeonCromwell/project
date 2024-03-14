package main

import (
	Config "example/auth-services/config"
	gintransport "example/auth-services/internal/module/auth/transport/gin"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("app.env")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	db, err := Config.Connect();
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	

	r := gin.Default()

	// protected route
	auth := r.Group("/auth")
	{
		auth.POST("/login", gintransport.LoginHandle(db))
		auth.POST("/register", gintransport.RegisterHandle(db))
		auth.POST("/vertify", gintransport.VertifyHandle(db))
		auth.GET("/refresh_token", gintransport.RefreshTokenHandle())
		auth.GET("/verify_token", gintransport.VerifyTokenHandle())
	}

	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}

