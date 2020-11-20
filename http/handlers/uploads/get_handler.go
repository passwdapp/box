package uploads

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GetHandler is a handler used to get the DB file
func GetHandler(ctx *fiber.Ctx) error {
	username := ctx.Locals("username").(string)

	return ctx.SendFile(fmt.Sprintf("./data/uploads/%s", username))
}
