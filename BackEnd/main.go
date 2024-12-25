package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
<<<<<<< HEAD
=======

>>>>>>> 7ba69d828ff760fbdf6839e9c755d67d51c5dd92
			log.Fatal(err)
		}
		utils.CheckUser(dbconn, customer.UserName)
		fmt.Print(customer)
		return c.JSON(customer)
	})

	app.Listen(":3000")
	defer dbconn.Close(context.Background())
}
