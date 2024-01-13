package auth

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/dev-newus/GoAlinBackend/src/handlers/auth/request"
	"github.com/dev-newus/GoAlinBackend/src/handlers/auth/response"
	"github.com/dev-newus/GoAlinBackend/src/utils"
)

func Register(ctx *fiber.Ctx, db *gorm.DB) error {
	var RegisterRequest request.Register
	var RegisterResponse response.Register

	if err := ctx.BodyParser(&RegisterRequest); err != nil {
		return utils.SendResponse(ctx, fiber.StatusBadRequest, RegisterRequest, utils.GetMessage("key"), map[string]string{})
	}

	// Validate the request
	if err := utils.Validate.Struct(&RegisterRequest); err != nil {
		transErrs := err.(validator.ValidationErrors).Translate(ctx.Locals("langTranslator").(ut.Translator))
		return utils.SendResponse(ctx, fiber.StatusBadRequest, &RegisterRequest, "", utils.GetError(RegisterRequest, transErrs))
	}

	RegisterResponse.Name = RegisterRequest.Name
	RegisterResponse.Email = RegisterRequest.Email

	return utils.SendResponse(ctx, 200, RegisterResponse, "", map[string]string{})
}
