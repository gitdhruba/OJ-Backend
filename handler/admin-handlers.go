/***********************************************************************
     Copyright (c) 2024 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package handler

import (
	db "oj-backend/database"
	"oj-backend/models"
	"oj-backend/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// login
func AdminLogin(c *fiber.Ctx) error {
	u := new(models.User)
	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	password := u.Password
	if res := db.DB.Where(&models.Admin{Username: u.Username}).Or(&models.Admin{Email: u.Username}).First(&u); (res.Error != nil) || (res.RowsAffected <= 0) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	accessToken, refreshToken, ok := util.GenerateTokens(strconv.Itoa(u.ID), true)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not generate tokens",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// create-admin
func CreateAdmin(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// create-contest
func CreateContest(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// create-contest
func ModifyContest(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// update-languages
func ModifyLanguages(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// add-question
func AddQuestion(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// update-question
func ModifyQuestion(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// add-testcase
func AddTestcase(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// delete-language
func DeleteLanguage(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// delete-contest
func DeleteContest(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// delete-question
func DeleteQuestion(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// delete-testcase
func DeleteTestcase(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// get-languages
func GetLanguages(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// get-contests
func GetContests(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// get-questions
func GetQuestions(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// get-question-details
func GetQuestionDetails(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// get-testcases
func GetTestcases(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}
