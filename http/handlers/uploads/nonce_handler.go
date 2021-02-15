package uploads

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/passwdapp/box/database"
	"github.com/passwdapp/box/models"
	"gorm.io/gorm"
)

// NonceHandler gives the latest upload nonce
func (h *Handler) NonceHandler(ctx *fiber.Ctx) error {
	username := ctx.Locals("username").(string)

	var upload models.Upload
	tx := database.GetDBConnection().Model(&models.Upload{}).Where("username = ?", username).First(&upload)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return ctx.SendStatus(404)
		}

		return ctx.SendStatus(500)
	}

	return ctx.JSON(models.NonceResponse{
		Nonce: upload.Nonce,
	})
}
