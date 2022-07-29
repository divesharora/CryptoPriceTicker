package controllers

import (
	"fmt"
	"net/mail"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/divesharora/KryptoBackendTask/api/db"
	"github.com/divesharora/KryptoBackendTask/api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.Users
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return c.Status(400).SendString("Invalid Email")
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

func Login(c *fiber.Ctx) error {
	var user models.Users
	var existingUser models.Users
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	userDB := db.GetDB()

	if err := userDB.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "User does not exist",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	token := &models.Token{
		UserID: existingUser.ID,
		Email:  existingUser.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, token).SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		return c.Status(500).SendString("Error Generating Token")
	}

	return c.JSON(fiber.Map{
		"message": "User logged in successfully",
		"token":   tokenString,
	})

}
