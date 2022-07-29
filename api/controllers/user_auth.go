package controllers

import (
	"fmt"

	"github.com/divesharora/KryptoBackendTask/api/db"
	"github.com/divesharora/KryptoBackendTask/api/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.Users
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString("Internal Server Error")
	}

	user.Password = string(pass)
	userDB := db.GetDB()
	createdUser := userDB.Create(&user)

	if createdUser.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "User already exists",
		})
	}
	return c.JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func Login()