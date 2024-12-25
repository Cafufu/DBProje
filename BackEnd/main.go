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
			// Hata oluşursa HTTP 400 dön
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON format",
			})
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
