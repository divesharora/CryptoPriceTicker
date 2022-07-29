package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/divesharora/KryptoBackendTask/api/cache"
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

	var allAlerts []models.Alert
	alertsDB.Find(&allAlerts).Where("email = ?", alert.Email)
	j, _ := json.Marshal(allAlerts)

	cache.SetValue(alert.Email, string(j), 0)
	return c.JSON(fiber.Map{
		"message": "Alert created successfully",
	})
}

func GetAlerts(c *fiber.Ctx) error {
	alertsDB := db.GetDB()
	alertsDB.AutoMigrate(&models.Alert{})
	email := c.Get("email")
	var alerts []models.Alert

	cv, _ := cache.GetValue(email)

	if cv != "" {
		json.Unmarshal([]byte(cv), &alerts)
	} else {
		alertsDB.Find(&alerts).Where("email = ?", email)
	}
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
	var a = models.Alert{}
	alertsDB.First(&a, ID)
	alertsDB.Delete(&a, "id = ?", ID)

	var allAlerts []models.Alert
	alertsDB.Find(&allAlerts).Where("email = ?", a.Email)
	j, _ := json.Marshal(allAlerts)

	cache.SetValue(a.Email, string(j), 0)
	return c.JSON(fiber.Map{
		"message": "Alert deleted successfully",
	})
}

func SendAlert(email string) {

	fmt.Println("Successfully sent mail to " + email)

}
