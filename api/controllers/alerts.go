package controllers

import (
	"fmt"

	"github.com/divesharora/KryptoBackendTask/api/db"
	"github.com/divesharora/KryptoBackendTask/api/models"
	"github.com/gofiber/fiber/v2"
)

func CreateAlert(c *fiber.Ctx) error {
	var alert models.Alert
	if err := c.BodyParser(&alert); err != nil {
		return err
	}

	alertsDB := db.GetDB()
	alertsDB.AutoMigrate(&models.Alert{})
	alert.Email = c.Locals("email").(string)
	createdAlert := alertsDB.Create(&alert)

	if createdAlert.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Alert already exists",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Alert created successfully",
	})
}

func GetAlerts(c *fiber.Ctx) error {
	alertsDB := db.GetDB()
	alertsDB.AutoMigrate(&models.Alert{})
	email := c.Get("email")
	var alerts []models.Alert
	alertsDB.Find(&alerts).Where("email = ?", email)
	return c.JSON(alerts)
}

func DeleteAlert(c *fiber.Ctx) error {
	alertsDB := db.GetDB()
	alertsDB.AutoMigrate(&models.Alert{})
	ID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}
	alertsDB.Delete(&models.Alert{}, "id = ?", ID)
	return c.JSON(fiber.Map{
		"message": "Alert deleted successfully",
	})
}

func SendAlert(email string) {

	fmt.Println("Successfully sent mail to " + email)

}
