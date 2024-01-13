package middlewares

import (
	"github.com/dev-newus/GoAlinBackend/src/utils"
	"github.com/gofiber/fiber/v2"
)

func SetupLanguage() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		lang := ctx.Get("lang")

		// Setup translation
		utils.SetupTranslation(lang)

		// Get translation object
		langTranslator := utils.GetTranslationFromHeader(lang)

		// Store translation object in locals for later use
		ctx.Locals("langTranslator", langTranslator)

		// Continue the request processing chain
		return ctx.Next()
	}
}
