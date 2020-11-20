package users

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/passwdapp/box/config"
	"github.com/passwdapp/box/database"
	"github.com/passwdapp/box/models"
	"golang.org/x/crypto/bcrypt"
)

// SignUpHandler is the handler used for signing up on the box
func SignUpHandler(ctx *fiber.Ctx) error {
	var currentUsers int64
	tx := database.GetDBConnection().Model(&models.User{}).Count(&currentUsers)
	if tx.Error != nil {
		return ctx.SendStatus(500)
	}

	maxUsers := config.GetConfig().MaxUsers
	if maxUsers <= currentUsers {
		return ctx.SendStatus(402)
	}

	body := new(models.SignUpBody)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.SendStatus(400)
	}

	if body.Username == "" || body.Password == "" {
		return ctx.SendStatus(400)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		log.Panicln(err)
	}

	tx = database.GetDBConnection().Create(&models.User{
		Username: body.Username,
		Password: string(hash),
	})

	if tx.Error != nil {
		return ctx.SendStatus(409)
	}

	return ctx.SendStatus(201)
}
