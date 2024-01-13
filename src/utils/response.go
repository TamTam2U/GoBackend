package utils

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  int               `json:"status"`
	Data    interface{}       `json:"data"`
	Message string            `json:"message"`
	Error   map[string]string `json:"error"`
}

func SendResponse(ctx *fiber.Ctx, status int, data interface{}, message string, err map[string]string) error {

	response := Response{
		Status:  status,
		Data:    data,
		Message: message,
		Error:   err,
	}

	statusCode := 200
	if status != 200 {
		statusCode = status
	}

	return ctx.Status(statusCode).JSON(response)
}
