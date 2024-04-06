package main

import (
	Config "example/auth-services/config"
	gintransport "example/auth-services/internal/module/auth/transport/gin"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
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
	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowedMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}, // Các HTTP methods được phép
		AllowedHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Access-Token", "Refresh-token"},
		ExposedHeaders:     []string{"Content-Length"},
		MaxAge: 86400, // Đảm bảo client không cần phải kiểm tra preflight thường xuyên
	})

	r.Use(func (ctx *gin.Context) {
		c.HandlerFunc(ctx.Writer, ctx.Request)
		ctx.Next()
		
	})

	// protected route
	auth := r.Group("/auth")
	{
		auth.POST("/login", gintransport.LoginHandle(db))
		auth.POST("/register", gintransport.RegisterHandle(db))
		auth.POST("/vertify", gintransport.VertifyHandle(db))
		auth.GET("/refresh_token", gintransport.RefreshTokenHandle())
		auth.GET("/verify_token", gintransport.VerifyTokenHandle())
		auth.GET("/get_user", gintransport.GetUserHandler(db))
		auth.POST("/update_user", gintransport.UpdateUserHandle(db))
	}

	r.Run(":5000") 
}

