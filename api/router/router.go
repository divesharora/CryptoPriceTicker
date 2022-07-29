package router

import (
	"github.com/divesharora/KryptoBackendTask/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func MountRoutes(c *fiber.App) {
	api := c.Group("api")

	{
	}

	comments := api.Group("users")

	{
		comments.Post("/signup/", controllers.CreateUser)
	}

	// edits := api.Group("edits")
	// {
	// 	edits.Post("/create/", controllers.CreateEdit)
	// 	edits.Get("/get/:version_id/:record_id", controllers.GetEdit)
	// 	edits.Get("/getAll/:record_id", controllers.GetAllEdits)
	// 	edits.Patch("/approve/:version_id/:record_id", controllers.ApproveEdit)
	// }
}
