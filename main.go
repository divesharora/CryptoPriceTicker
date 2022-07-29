package main

import (
	"fmt"

	"github.com/divesharora/KryptoBackendTask/api/db"
	"github.com/divesharora/KryptoBackendTask/api/router"
	"github.com/divesharora/KryptoBackendTask/api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	utils.ImportEnv()

	app := fiber.New(fiber.Config{})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	},
	)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "",
		AllowHeaders: "*",
	}))

	db.Connect()

	router.MountRoutes(app)

	port := utils.GetPort()

	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

}
