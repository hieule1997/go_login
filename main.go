package main

import (
	"github.com/gin-gonic/gin"
	"login_jwt/controllers"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.POST("/login", controllers.Singin)
	v1.POST("/register",controllers.Register)
	router.Run()
}


