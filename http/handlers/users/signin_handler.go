package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/passwdapp/box/database"
	"github.com/passwdapp/box/models"
	"github.com/passwdapp/box/utils"
	"golang.org/x/crypto/bcrypt"
)

// SignInHandler is the handler used to sign into box
func SignInHandler(ctx *fiber.Ctx) error {
	body := new(models.SignInBody)

	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	if body.Username == "" || body.Password == "" {
		return ctx.SendStatus(400)
	}

	var user models.User
	tx := database.GetDBConnection().Model(&models.User{}).Where("username = ?", string(body.Username)).First(&user)

	if tx.Error != nil {
		return ctx.SendStatus(401)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return ctx.SendStatus(401)
	}

	accessToken, refreshToken, err := utils.GenerateLoginTokens(user)
	if err != nil {
		return ctx.SendStatus(500)
	}

	response := models.SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return ctx.JSON(response)
}
