/***********************************************************************
     Copyright (c) 2025 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package handler

import (
	db "oj-backend/database"
	"oj-backend/models"
	"oj-backend/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// signup
func Signup(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	// check if user already exists either by email or username
	if res := db.DB.Where(&models.User{Email: user.Email}).Or(&models.User{Username: user.Username}).First(&models.User{}); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	} else if res.RowsAffected > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "User already exists",
		})
	}

	// validate email and password
	if check := util.IsValidEmail(user.Email); !check {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid email",
		})
	}

	if check := util.IsValidPassword(user.Password); !check {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "weak password",
		})
	}

	// hash password
	if !util.HashPassword(&user.Password) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not hash password",
		})
	}

	// create user
	if res := db.DB.Create(&user); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create user",
		})
	}

	// generate access token and refresh token
	accessToken, refreshToken, ok := util.GenerateTokens(strconv.Itoa(user.ID), false)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not generate tokens",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":       "User created successfully",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// login
func Login(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	// get user password(hashed) from db
	var password string
	if res := db.DB.Where(&models.User{Email: user.Email, Username: user.Username}).Pluck("password", &password); res.Error != nil || res.RowsAffected <= 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// compare password
	if !util.ComparePassword(password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// generate access token and refresh token
	accessToken, refreshToken, ok := util.GenerateTokens(strconv.Itoa(user.ID), false)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not generate tokens",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "Login successful",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// reclaim access token
func ReclaimAccessToken(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// logout
func Logout(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}
