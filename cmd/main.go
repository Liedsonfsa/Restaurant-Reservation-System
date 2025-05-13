package main

import (
	"github.com/Liedsonfsa/Restaurant-Reservation-System/internal/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.POST("/login", controllers.Login)
	server.POST("/register", controllers.Register)

	server.Run()
}