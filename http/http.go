package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/passwdapp/box/config"
	"github.com/passwdapp/box/http/handlers"
	"github.com/passwdapp/box/http/middleware"
)

// InitHTTP initializes the HTTP server
func InitHTTP() {
	conf := config.GetConfig()

	app := fiber.New(fiber.Config{
		ServerHeader: fmt.Sprintf("passwd_box/fiber/%s", config.Version),
		Prefork:      false,
	})

	secretKeyMiddleware := middleware.SecretKeyMiddleware{}
	secretKeyMiddleware.InitMiddleware(conf)

	app.Use(recover.New())
	app.Use(secretKeyMiddleware.Handler)
	app.Use(logger.New())

	v1Group := app.Group("/v1")
	usersGroup := v1Group.Group("/users")

	usersGroup.Post("/signup", handlers.SignUpHandler)
	usersGroup.Post("/signin", handlers.SignInHandler)
	usersGroup.Post("/refresh", handlers.RefreshHandler)

	app.Listen(conf.ListenAddress)
}
