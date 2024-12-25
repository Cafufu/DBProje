package main

import (
	"encoding/json"
	"fmt"
	"mymodule/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	dbconn := utils.DbConnect()
	app.Post("/register", func(c *fiber.Ctx) error {
		body := c.Body()
		var customer utils.Customer
		err := json.Unmarshal(body, &customer)
		if err != nil {
			// Hata oluşursa HTTP 400 dön
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON format",
			})
		}
		utils.CheckUser(dbconn, customer.UserName)
		fmt.Print(customer)
		return c.JSON(customer)
	})

	app.Listen(":3000")

}
