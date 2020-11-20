package uploads

import "github.com/gofiber/fiber/v2"

// NonceHandler gives the latest upload nonce
func NonceHandler(ctx *fiber.Ctx) error {
	return ctx.SendStatus(200)
}
