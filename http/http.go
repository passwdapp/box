package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/passwdapp/box/config"
	"github.com/passwdapp/box/http/handlers"
	"github.com/passwdapp/box/http/handlers/uploads"
	"github.com/passwdapp/box/http/handlers/users"
	"github.com/passwdapp/box/http/middleware"
)

// InitHTTP initializes the HTTP server
func InitHTTP() {
	conf := config.GetConfig()

	app := fiber.New(fiber.Config{
		ServerHeader: fmt.Sprintf("passwd_box/fiber/%s", config.Version),
		Prefork:      false,
		BodyLimit:    256 * 1024, // 256 kb
	})

	secretKeyMiddleware := middleware.SecretKeyMiddleware{}
	secretKeyMiddleware.InitMiddleware(conf)

	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(secretKeyMiddleware.Handler)
	app.Use(logger.New())

	v1Group := app.Group("/v1")
	v1Group.Get("/ping", handlers.PingHandler)

	usersGroup := v1Group.Group("/users")
	usersGroup.Post("/signup", users.SignUpHandler)
	usersGroup.Post("/signin", users.SignInHandler)
	usersGroup.Post("/refresh", users.RefreshHandler)

	protectedGroup := v1Group.Group("/protected")
	protectedGroup.Use(jwtware.New(jwtware.Config{
		SigningKey:    []byte(config.GetConfig().JWTSecret),
		SigningMethod: "HS512",
	}))
	protectedGroup.Use(middleware.UsernameMiddleware)

	protectedGroup.Get("/ping", handlers.PingHandler)

	uploadsGroup := protectedGroup.Group("/uploads")
	uploadsGroup.Get("/nonce", uploads.NonceHandler)
	uploadsGroup.Post("/new", uploads.UploadHandler)
	uploadsGroup.Get("/get", uploads.GetHandler)

	app.Listen(conf.ListenAddress)
}
