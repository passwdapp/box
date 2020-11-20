package middleware

import (
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/passwdapp/box/database"
	"github.com/passwdapp/box/models"
)

// UsernameMiddleware parses the incoming JWT and stores the username
func UsernameMiddleware(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	username := claims["username"].(string)
	ctx.Locals("username", username)

	if username == "" {
		return ctx.SendStatus(401)
	}

	var userRecord models.User
	tx := database.GetDBConnection().Model(&models.User{}).Where("username = ?", string(username)).First(&userRecord)
	if tx.Error != nil {
		return ctx.SendStatus(401)
	}

	return ctx.Next()
}
