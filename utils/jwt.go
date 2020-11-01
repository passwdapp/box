package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/passwdapp/box/config"
	"github.com/passwdapp/box/database"
	"github.com/passwdapp/box/models"
)

// GenerateLoginTokens creates a JWT with 1 hour validity and a refresh token (which is synced with the DB)
// Return (error, accessToken, )
func GenerateLoginTokens(user models.User) (at, rt string, err error) {
	if user.Username == "" {
		return "", "", errors.New("Empty username")
	}

	jwtSecret := config.GetConfig().JWTSecret

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().UTC().Add(time.Hour).Unix()
	claims["username"] = user.Username

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	accessTokenSigned, err := accessToken.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateRandomStringURLSafe(64)

	if err != nil {
		return "", "", err
	}

	tx := database.GetDBConnection().Create(&models.RefreshToken{
		Token:    refreshToken,
		Username: user.Username,
	})

	if tx.Error != nil {
		return "", "", tx.Error
	}

	return accessTokenSigned, refreshToken, nil
}
