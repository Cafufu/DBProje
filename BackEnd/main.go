package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Customer struct {
	abone_num int
	price     int
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

	app.Get("/furkan", func(c *fiber.Ctx) error {
		fmt.Println("Request Furkana geldi")
		return c.SendString("Furkan")
	})

	app.Listen(":3000")

}
