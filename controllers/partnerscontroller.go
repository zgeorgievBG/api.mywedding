package controllers

import (
	"fmt"
	"net/http"

	"api.mywedding/auth"
	"api.mywedding/database"
	"api.mywedding/models"
	"github.com/gin-gonic/gin"
)

func GetAllPartnersCards(context *gin.Context) {
	type result struct {
		UserID         string
		Email          string
		FirstName      string
		LastName       string
		Price          string
		Description    string
		ProfilePicture string
		ArtistType     string
	}

	var result1 []result

	database.DB.Model(&models.User{}).Select("users.user_id,users.email,users.first_name, users.last_name, artists.price, artists.profile_picture, artists.description, artists.artist_type, artists.operating_in").Joins("left join artists on artists.artist_id = users.user_id").Scan(&result1)
	context.JSON(http.StatusCreated, gin.H{"users": result1})

}

func CreatePartnerCard(context *gin.Context) {
	tokenString := context.GetHeader("Authorization")
	_, decodedUserId, _ := auth.ValidateToken(tokenString)
	var artist models.Artist
	record := database.DB.Where("artist_id = ?", decodedUserId).First(&artist)

	if record.Error == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User already created a card"})
		context.Abort()
		return
	}

	if err := context.ShouldBindJSON(&artist); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	artist.ArtistID = decodedUserId
	result := database.DB.Create(&artist)
	if result.Error != nil {
		fmt.Println("Error creating artist card")
	}
	context.JSON(http.StatusCreated, gin.H{"id": artist.ID, "artist_id": artist.ArtistID})
}

func UpdatePartnerCard(context *gin.Context) {
	cardId, _ := context.Params.Get("id")
	if cardId == "" {
		fmt.Println("No card id provided")
	}
	tokenString := context.GetHeader("Authorization")
	_, decodedUserId, _ := auth.ValidateToken(tokenString)
	var data models.Artist
	var artist models.Artist

	record := database.DB.Where("id = ?", cardId).First(&artist)

	if record.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "There is no artist card for this user"})
		context.Abort()
		return
	}
	fmt.Println(decodedUserId, artist.ArtistID)
	if artist.ArtistID != decodedUserId {
		context.JSON(http.StatusBadRequest, gin.H{"error": "You are not allowed to update this card"})
		context.Abort()
		return
	}
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	database.DB.Model(&artist).Where("id = ?", cardId).Updates(models.Artist{Price: data.Price, ProfilePicture: data.ProfilePicture, Description: data.Description, ArtistType: data.ArtistType, OperatingIn: data.OperatingIn, Instagram: data.Instagram, Facebook: data.Facebook, Website: data.Website})

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

	database.DB.Where("user_id = ?", data.ArtistID).Delete(&artist)

	context.JSON(http.StatusOK, gin.H{"message": "Artist deleted successfully"})
}
