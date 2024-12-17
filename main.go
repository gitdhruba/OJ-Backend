/***********************************************************************
     Copyright (c) 2024 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package main

import (
	"log"
	"oj-backend/config"

	"github.com/gofiber/fiber/v2"
)

const (
	port = ":3000"
)

func main() {

	// load environment variables
	config.LoadEnv()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		log.Println("GET")
		return c.SendStatus(fiber.StatusOK)
	})

	log.Fatal(app.Listen(port))
}
