package models

import "gorm.io/gorm"

// RefreshToken is the model for the refresh_token generate and stored into the DB
// It is used to refresh the JWTs. The JWTs have a validity of 1h
type RefreshToken struct {
	gorm.Model
	Token string `json:"refresh_token" gorm:"uniqueIndex"`
	User  User
}
