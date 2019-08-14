package main

import (
	"github.com/gin-gonic/gin"
	"login_jwt/controllers"
	"login_jwt/middleware"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.POST("/register",controllers.Register)
	v1.POST("/login", controllers.Singin)
	v1.Use(middleware.Jwt_required())
	{
		v1.PATCH("/user", controllers.UpdateUser)
	}
	router.Run()
}