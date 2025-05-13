package main

import (
	"github.com/Liedsonfsa/Restaurant-Reservation-System/internal/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	user := server.Group("/usuarios")
	{
		user.POST("/login", controllers.Login)
		user.POST("/registrar", controllers.Register)
	}

	tables := server.Group("/mesas")
	{
		tables.GET("/", nil)
		tables.POST("/", nil)
		tables.PUT("/:id", nil)
		tables.DELETE("/:id", nil)
	}

	reservations := server.Group("/reservas")
	{
		reservations.POST("/", nil)
		reservations.GET("/", nil)
		reservations.PUT("/:id/cancelar", nil)
	}

	server.Run()
}