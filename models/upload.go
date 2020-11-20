package models

import "gorm.io/gorm"

// Upload defines the model for upload stored into the DB
type Upload struct {
	gorm.Model
	Nonce    string `json:"nonce" gorm:"uniqueIndex"`
	Username string `json:"username" gorm:"not null"`
}

// NonceResponse is the result for a nonce GET request
type NonceResponse struct {
	Nonce string `json:"nonce"`
}
