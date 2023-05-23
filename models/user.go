package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID      string    `db:"user_id"`
	Email       string    `db:"email" gorm:"unique; not null"`
	Password    string    `db:"password" gorm:"not null"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `db:"date_of_birth" gorm:"not null"`
	Gender      string    `db:"gender" gorm:"not null"`
	Address     string    `db:"address" gorm:"not null"`
	PhoneNumber string    `db:"phone_number" gorm:"not null"`
}

type Artist struct {
	gorm.Model
	ArtistID       string `db:"artist_id" gorm:"foreignKey:UserID"`
	Price          int    `db:"price" gorm:"not null"`
	ProfilePicture string `json:"profile_picture"`
	Description    string `db:"description"`
	ArtistType     string `json:"artist_type" gorm:"not null"`
	OperatingIn    string `json:"operating_in" gorm:"not null"`
	Instagram      string `db:"instagram"`
	Facebook       string `db:"facebook"`
	Website        string `db:"website"`
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
