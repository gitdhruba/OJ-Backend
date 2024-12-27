/***********************************************************************
     Copyright (c) 2024 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package handler

import "github.com/gofiber/fiber/v2"

// signup
func Signup(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// login
func Login(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// reclaim access token
func ReclaimAccessToken(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// logout
func Logout(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}
