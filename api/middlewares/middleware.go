package middlewares

import (
	"fmt"
	// "strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/divesharora/KryptoBackendTask/api/models"
	"github.com/gofiber/fiber/v2"
)

func VerifyJWT(c *fiber.Ctx) error {
	var header = c.Get("auth-token")
	fmt.Println(header)

	// header = strings.TrimSpace(header)

	if header == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	tk := &models.Token{}

	_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	fmt.Println(err)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}
