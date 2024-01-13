package routes

import (
	"github.com/dev-newus/GoAlinBackend/src/handlers/agency"
	"github.com/dev-newus/GoAlinBackend/src/handlers/auth"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RoutesInit(app *fiber.App, db *gorm.DB) {
	agencyRoute := app.Group("/agency")
	agencyRoute.Get("/GetAll", func(ctx *fiber.Ctx) error {
		return agency.GetAll(ctx, db)
	})

	authRoute := app.Group("/auth")
	authRoute.Post("/register", func(ctx *fiber.Ctx) error {
		return auth.Register(ctx, db)
	})
}
