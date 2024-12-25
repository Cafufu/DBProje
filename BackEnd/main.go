package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Customer struct {
	name        string `json:"name"`
	surname     string `json:"surname"`
	userName    string `json:"userName"`
	phoneNumber string `json:"phoneNumber"`
	password    string `json:"password"`
}
type PostData struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

func main() {

	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		fmt.Println("Request geldi")
		body := c.Body()
		var data PostData
		err := json.Unmarshal(body, &data)
		if err != nil {
			fmt.Printf("err was %v", err)
		}
		data.UserID = data.UserID + 10
		fmt.Println(data)
		return c.JSON(data)
	})

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

	app.Get("/furkan", func(c *fiber.Ctx) error {
		fmt.Println("Request Furkana geldi")
		return c.SendString("Furkan")
	})

	app.Listen(":3000")

}
