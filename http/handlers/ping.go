package handlers

import "github.com/gofiber/fiber/v2"

// PingHandler is retuns a plain response "OK" if all the middleware based checks pass
func PingHandler(ctx *fiber.Ctx) error {
	return ctx.Send("OK")
}
