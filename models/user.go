package models

import "gorm.io/gorm"

// User is the model for the users created in the DB
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex"`
	Password string `json:"password" gorm:"uniqueIndex"`
}

// SignUpBody contains the body for the sign-up request
type SignUpBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignInBody contains the body for the sign-up request
type SignInBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
