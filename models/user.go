package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID      string    `json:"user_id"`
	Email       string    `json:"email" gorm:"unique"`
	Password    string    `json:"password"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
}

type Artist struct {
	gorm.Model
	UserID         string `json:"user_id"`
	User           User   `gorm:"foreignKey:UserID"`
	Price          int    `json:"price"`
	ProfilePicture string `json:"profile_picture"`
	Description    string `json:"description"`
	ArtistType     string `json:"artist_type"`
	OperatingIn    string `json:"operating_in"`
	Instagram      string `json:"instagram"`
	Facebook       string `json:"facebook"`
	Website        string `json:"website"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
