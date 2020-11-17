package models

import "gorm.io/gorm"

// RefreshToken is the model for the refresh_token generate and stored into the DB
// It is used to refresh the JWTs. The JWTs have a validity of 1h
type RefreshToken struct {
	gorm.Model
	Token    string `json:"refresh_token" gorm:"uniqueIndex"`
	Username string `gorm:"not null"`
}

// SignInResponse contains the response for a sign-in request
type SignInResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// RefreshResponse is the response for a token refreh request
type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}
