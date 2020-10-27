package models

import "gorm.io/gorm"

// User is the model for the users created in the DB
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex"`
	Password string `json:"password" gorm:"uniqueIndex"`
}
