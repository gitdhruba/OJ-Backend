/***********************************************************************
     Copyright (c) 2024 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package main

import (
	"fmt"
	"oj-backend/config"
	db "oj-backend/database"
	"oj-backend/router"

	"github.com/gofiber/fiber/v2"
)

const (
	port = ":3000"
)

func main() {

	// load environment variables
	config.LoadEnv()

	// connect to database
	db.ConnectDB()

	// create new fiber instance
	app := fiber.New()

	// setup routes
	router.SetupRoutes(app)

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	log.Println("GET")
	// 	return c.SendStatus(fiber.StatusOK)
	// })

	fmt.Println(app.Listen(port))
}
