package uploads

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/passwdapp/box/database"
	"github.com/passwdapp/box/models"
	"github.com/passwdapp/box/utils"
	"gorm.io/gorm"
)

// UploadHandler handles the upload of the DB and then returns a nonce
func UploadHandler(ctx *fiber.Ctx) error {
	username := ctx.Locals("username").(string)

	file, err := ctx.FormFile("db")
	if err != nil {
		return ctx.SendStatus(400)
	}

	err = ctx.SaveFile(file, fmt.Sprintf("./data/uploads/%s", username))
	if err != nil {
		return ctx.SendStatus(500)
	}

	var upload models.Upload
	tx := database.GetDBConnection().Model(&models.Upload{}).Where("username = ?", username).First(&upload)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			nonce, err := utils.GenerateRandomString(24)

			if err != nil {
				return ctx.SendStatus(500)
			}

			tx = database.GetDBConnection().Create(&models.Upload{
				Username: username,
				Nonce:    nonce,
			})
			if tx.Error != nil {
				return ctx.SendStatus(500)
			}

			return ctx.JSON(models.NonceResponse{
				Nonce: nonce,
			})
		}

		return ctx.SendStatus(500)
	}

	upload.Nonce, err = utils.GenerateRandomString(24)
	if err != nil {
		ctx.SendStatus(500)
	}

	tx = database.GetDBConnection().Save(&upload)
	if tx.Error != nil {
		return ctx.SendStatus(500)
	}

	return ctx.JSON(models.NonceResponse{
		Nonce: upload.Nonce,
	})
}
