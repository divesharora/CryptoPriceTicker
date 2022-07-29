package middlewares

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/divesharora/KryptoBackendTask/api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func VerifyJWT(c *fiber.Ctx) error {
	var header = c.Get("auth-token")

	if header == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	tk := &models.Token{}

	_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("JWT_SECRET")), nil
	})
	fmt.Println(tk.Email)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	c.Locals("email", tk.Email)

	return c.Next()
}
