package uploads

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/passwdapp/box/database"
	"github.com/passwdapp/box/models"
	"gorm.io/gorm"
)

// UploadHandler handles the upload of the DB and then returns a nonce
func (h *Handler) UploadHandler(ctx *fiber.Ctx) error {
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
			nonce := "0"

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

	oldNonce, err := strconv.Atoi(upload.Nonce)
	if err != nil {
		upload.Nonce = "0"
	}

	upload.Nonce = fmt.Sprintf("%d", oldNonce+1)

	tx = database.GetDBConnection().Save(&upload)
	if tx.Error != nil {
		return ctx.SendStatus(500)
	}

	return ctx.JSON(models.NonceResponse{
		Nonce: upload.Nonce,
	})
}
