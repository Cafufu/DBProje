package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Customer struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	UserName    string `json:"userName"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

func main() {

	app := fiber.New()

	app.Post("/register", func(c *fiber.Ctx) error {
		body := c.Body()
		var customer Customer
		err := json.Unmarshal(body, &customer)
		if err != nil {
			fmt.Printf("err was %v", err)
		}
		fmt.Print(customer)
		return c.JSON(customer)
	})

	app.Listen(":3000")

}
