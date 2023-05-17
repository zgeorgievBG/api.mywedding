package controllers

import (
	"net/http"

	"api.mywedding/database"
	"api.mywedding/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	var existingUser models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	result := database.DB.First(&existingUser, "email = ?", user.Email)

	if result.Error != nil {
		userID := uuid.NewV4().String()
		user.UserID = userID
		record := database.DB.Create(&user)
		if record.Error != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
			context.Abort()
			return
		}
		context.JSON(http.StatusCreated, gin.H{"id": user.ID, "email": user.Email, "userId": user.UserID})
	} else {
		context.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
		context.Abort()
		return
	}

}

func LoginUser(context *gin.Context) {
	context.JSON(200, gin.H{"status": "OK", "message": "LoginUser"})
}
