package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/passwdapp/box/database"
	"github.com/passwdapp/box/models"
	"github.com/passwdapp/box/utils"
)

// RefreshHandler is used to refresh the access tokens
func RefreshHandler(ctx *fiber.Ctx) error {
	body := new(models.RefreshBody)

	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	if body.RefreshToken == "" {
		return ctx.SendStatus(400)
	}

	valid, username, err := utils.VerifyRefreshToken(body.RefreshToken)

	if !valid {
		return ctx.SendStatus(401)
	}
	if err != nil {
		return ctx.SendStatus(500)
	}

	var user models.User
	tx := database.GetDBConnection().Model(&models.User{}).Where("username = ?", string(username)).First(&user)
	if tx.Error != nil {
		return ctx.SendStatus(401)
	}

	refreshedAccessToken, err := utils.GenerateJWT(user)
	if err != nil {
		return err
	}

	return ctx.JSON(models.RefreshResponse{
		AccessToken: refreshedAccessToken,
	})
}
