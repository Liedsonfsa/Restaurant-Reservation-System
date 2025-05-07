package controllers

import (
	"net/http"

	"github.com/Liedsonfsa/Restaurant-Reservation-System/internal/models"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"usuário": user})
}