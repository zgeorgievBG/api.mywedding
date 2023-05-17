package controllers

import (
	"fmt"
	"net/http"

	"api.mywedding/database"
	"api.mywedding/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPartners(context *gin.Context) {
	var users []models.User
	var artists []models.Artist

	database.DB.Joins("JOIN artists ON artists.user_id = users.user_id").Find(&users).Find(&artists)
	fmt.Printf("%+v\n", artists)
}

func CreatePartner(context *gin.Context) {
	var artist models.Artist
	if err := context.ShouldBindJSON(&artist); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.DB.Having("user_id = ?", artist.UserID)
	if record.Error != nil {
		fmt.Println("There is no such user")
	}

	result := database.DB.Create(&artist)
	if result.Error != nil {
		fmt.Println("Error creating artist")
	}

	context.JSON(http.StatusCreated, gin.H{"id": artist.ID, "user_id": artist.UserID})

}

func UpdatePartner(context *gin.Context) {
	var data models.Artist
	var artist models.Artist
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	database.DB.Model(&artist).Where("user_id = ?", data.UserID).Updates(models.Artist{Price: data.Price, ProfilePicture: data.ProfilePicture, Description: data.Description, ArtistType: data.ArtistType, OperatingIn: data.OperatingIn, Instagram: data.Instagram, Facebook: data.Facebook, Website: data.Website})

	context.JSON(http.StatusOK, gin.H{"message": "Artist updated successfully"})
}

func DeletePartner(context *gin.Context) {
	var data models.Artist
	var artist models.Artist
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	database.DB.Where("user_id = ?", data.UserID).Delete(&artist)

	context.JSON(http.StatusOK, gin.H{"message": "Artist deleted successfully"})
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController(DB *gorm.DB) PostController {
	return PostController{DB}
}
