package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/passwdapp/box/config"
	"github.com/passwdapp/box/database"
	"github.com/passwdapp/box/models"
	"github.com/passwdapp/box/utils"
)

// SignUpHandler is the handler used for signing up on the box
func SignUpHandler(ctx *fiber.Ctx) error {
	body := new(models.SignUpBody)

	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	if body.Username == "" || body.Password == "" {
		return ctx.SendStatus(400)
	}

	config := &config.PasswordConfig{
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	}

	hash, err := utils.GeneratePassword(config, body.Password)
	if err != nil {
		log.Panicln(err)
	}

	database.GetDBConnection().Create(&models.User{
		Username: body.Username,
		Password: hash,
	})

	return ctx.SendStatus(201)
}
