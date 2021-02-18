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
func InitHTTP(cfg *config.Config) {
	app := fiber.New(fiber.Config{
		ServerHeader: fmt.Sprintf("passwd_box/fiber/%s", config.Version),
		Prefork:      false,
		BodyLimit:    256 * 1024, // 256 kb
	})

	app.Use(logger.New())

	secretKeyMiddleware := middleware.SecretKeyMiddleware{}
	secretKeyMiddleware.InitMiddleware(cfg)

	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(secretKeyMiddleware.Handler)

	v1Group := app.Group("/v1")
	v1Group.Get("/ping", handlers.PingHandler)

	usersGroup := v1Group.Group("/users")

	userHandlers := users.Handler{}
	userHandlers.Init(cfg)

	usersGroup.Post("/signup", userHandlers.SignUpHandler)
	usersGroup.Post("/signin", userHandlers.SignInHandler)
	usersGroup.Post("/refresh", userHandlers.RefreshHandler)

	protectedGroup := v1Group.Group("/protected")
	protectedGroup.Use(jwtware.New(jwtware.Config{
		SigningKey:    []byte(cfg.JWTSecret),
		SigningMethod: "HS512",
	}))
	protectedGroup.Use(middleware.UsernameMiddleware)

	protectedGroup.Get("/ping", handlers.PingHandler)

	uploadsGroup := protectedGroup.Group("/uploads")

	uploadsHandlers := uploads.Handler{}
	uploadsHandlers.Init(cfg)

	uploadsGroup.Get("/nonce", uploadsHandlers.NonceHandler)
	uploadsGroup.Post("/new", uploadsHandlers.UploadHandler)
	uploadsGroup.Get("/get", uploadsHandlers.GetHandler)

	app.Listen(cfg.ListenAddress)
}
