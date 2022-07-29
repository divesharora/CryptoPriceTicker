package router

import (
	"github.com/divesharora/KryptoBackendTask/api/controllers"
	"github.com/divesharora/KryptoBackendTask/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func MountRoutes(c *fiber.App) {
	api := c.Group("api")

	{
	}

	auth := api.Group("users")

	{
		auth.Post("/signup/", middlewares.VerifyJWT, controllers.CreateUser)
		auth.Post("/login/", controllers.Login)
	}

	alerts := api.Group("alerts")
	{
		alerts.Post("/", middlewares.VerifyJWT, controllers.CreateAlert)
		alerts.Get("/", middlewares.VerifyJWT, controllers.GetAlerts)
		alerts.Delete("/:id", middlewares.VerifyJWT, controllers.DeleteAlert)
	}

}
