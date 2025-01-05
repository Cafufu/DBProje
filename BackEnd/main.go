package main

import (
	"encoding/json"
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
		return c.JSON(retValue)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		body := c.Body()
		var login utils.LoginInput
		err := json.Unmarshal(body, &login)
		if err != nil {
			log.Fatal(err)
		}
		id := utils.CheckLogin(dbconn, login)

		return c.JSON(id)
	})

	app.Post("/insert", func(c *fiber.Ctx) error {
		body := c.Body()
		var bill utils.Bill
		var retVal int
		err := json.Unmarshal(body, &bill)
		if err != nil {
			log.Fatal(err)
		}
		exist := utils.CheckBill(dbconn, bill)

		if exist {
			retVal = utils.UpdateBill(dbconn, bill)
		} else {
			retVal = utils.InsertBill(dbconn, bill)
		}

		return c.JSON(retVal)
	})

	app.Post("/show", func(c *fiber.Ctx) error {
		body := c.Body()
		var billInfo utils.BillInfo
		err := json.Unmarshal(body, &billInfo)
		if err != nil {
			log.Fatal(err)
		}
		Bills := utils.ShowBills(dbconn, billInfo)
		if len(Bills) == 0 {
			return c.JSON(-1)
		}
		return c.JSON(Bills)
	})
	app.Post("/carbon", func(c *fiber.Ctx) error {
		body := c.Body()
		var userId int
		err := json.Unmarshal(body, &userId) // int olarak alıyorum burada ama furkan string gonderirse convert yapcaz.
		if err != nil {
			log.Fatal(err)
		}

		carbonFootprint := utils.ShowCarbonFootPrint(dbconn, userId) // string olarak gönderiyorum duruma göre değişebiliriz.
		return c.JSON(carbonFootprint)
	})
	app.Post("/remove", func(c *fiber.Ctx) error {
		body := c.Body()
		var bill utils.Bill
		err := json.Unmarshal(body, &bill)
		if err != nil {
			log.Fatal(err)
		}

		retVal := utils.DeleteBill(dbconn, bill)
		if retVal == 1 {
			utils.UpdateCarbonFootPrint(dbconn, bill.UserId) // her fatura delete edildiğinde karbon ayakizi update ediliyor
		}
		return c.JSON(retVal)
	})
	app.Post("/analiz", func(c *fiber.Ctx) error {
		body := c.Body()
		var userId int
		err := json.Unmarshal(body, &userId) // int olarak alıyorum burada ama furkan string gonderirse convert yapcaz.
		if err != nil {
			log.Fatal(err)
		}
		analizString := utils.Analiz(dbconn, userId) // string olarak gönderiyorum duruma göre değişebiliriz.
		return c.JSON(analizString)
	})

	app.Listen(":3000")

}
