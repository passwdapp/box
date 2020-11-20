package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// UsernameMiddleware parses the incoming JWT and stores the username
func UsernameMiddleware(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	username := claims["username"].(string)
	ctx.Locals("username", username)

	return ctx.Next()
}
