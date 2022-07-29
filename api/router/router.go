package router

import (
	"github.com/divesharora/KryptoBackendTask/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func MountRoutes(c *fiber.App) {
	api := c.Group("api")

	{
	}

	auth := api.Group("users")

	{
		auth.Post("/signup/", controllers.CreateUser)
		auth.Post("/login/", controllers.Login)
	}

}
