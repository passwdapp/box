package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/passwdapp/box/config"
	"github.com/passwdapp/box/http/middleware"
)

// InitHTTP initializes the HTTP server
func InitHTTP() {
	conf := config.GetConfig()

	app := fiber.New(fiber.Config{
		ServerHeader: fmt.Sprintf("passwd_box/fiber/%s", config.Version),
	})

	secretKeyMiddleware := middleware.SecretKeyMiddleware{}
	secretKeyMiddleware.InitMiddleware(conf)

	app.Use(recover.New())
	app.Use(secretKeyMiddleware.Handler)
	app.Use(logger.New())

	app.Listen(conf.ListenAddress)
}
