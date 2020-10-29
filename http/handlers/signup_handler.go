package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/passwdapp/box/database"
	"github.com/passwdapp/box/models"
	"golang.org/x/crypto/bcrypt"
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

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		log.Panicln(err)
	}

	tx := database.GetDBConnection().Create(&models.User{
		Username: body.Username,
		Password: string(hash),
	})

	if tx.Error != nil {
		return ctx.SendStatus(409)
	}

	return ctx.SendStatus(201)
}
