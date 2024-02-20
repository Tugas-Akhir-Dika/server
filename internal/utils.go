package internal

import "github.com/gofiber/fiber/v2"

const AUTH_KEY = "lucukamudek"

func GetAuthStatus(ctx *fiber.Ctx) bool {
	token := ctx.GetRespHeader("Authorization")
	return token == AUTH_KEY
}
