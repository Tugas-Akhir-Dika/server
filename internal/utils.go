package internal

import (
	"github.com/gofiber/fiber/v2"
)

const AUTH_KEY = "lucukamudek"

func IsAuth(ctx *fiber.Ctx) bool {
	token := ctx.Get("Authorization")
	return token == AUTH_KEY
}
