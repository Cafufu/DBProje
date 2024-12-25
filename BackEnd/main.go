package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mymodule/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Accept, Origin")
		if c.Method() == fiber.MethodOptions {
			return c.SendStatus(fiber.StatusOK)
		}
		return c.Next()
	})

	dbconn := utils.DbConnect()

	app.Post("/register", func(c *fiber.Ctx) error {
		body := c.Body()
		var customer utils.Customer
		err := json.Unmarshal(body, &customer)
		if err != nil {
			log.Fatal(err)
		}
		retValue := utils.CheckUser(dbconn, customer.UserName)
		if retValue == 1 {
			utils.Insert(dbconn, customer)
		}
		fmt.Print(retValue)
		return c.JSON(retValue)
	})

	app.Listen(":3000")

}
