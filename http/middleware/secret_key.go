package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/passwdapp/box/config"
)

// SecretKeyMiddleware is the middleware to reject any unauthorized request
type SecretKeyMiddleware struct {
	conf *config.Config
}

// InitMiddleware initializes the middleware
func (m *SecretKeyMiddleware) InitMiddleware(conf *config.Config) {
	m.conf = conf
}

// Handler is the middleware handler for fiber
func (m *SecretKeyMiddleware) Handler(ctx *fiber.Ctx) error {
	secretKey := ctx.Request().Header.Peek("X-Secret-Key")

	if secretKey == nil || string(secretKey) != m.conf.SecretKey {
		return ctx.SendStatus(401)
	}

	return ctx.Next()
}
