package main

import (
	"github.com/dev-newus/GoAlinBackend/src/middlewares"
	"github.com/dev-newus/GoAlinBackend/src/routes"
	"github.com/dev-newus/GoAlinDatabase"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use(middlewares.SetupLanguage())
	dbconf := GoAlinDatabase.Config{
		Host:               "localhost",
		Port:               "3306",
		User:               "root",
		Password:           "",
		Name:               "base",
		MaxIdleConnections: 10,
		MaxOpenConnections: 100,
	}
	DB := GoAlinDatabase.GetTenantDB("BASE", dbconf)
	//initial route
	routes.RoutesInit(app, DB)
	app.Listen(":3000")
}
