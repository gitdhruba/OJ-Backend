/***********************************************************************
     Copyright (c) 2024 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package handler

import "github.com/gofiber/fiber/v2"

// get contest-list
func GetContestList(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// get question-list
func GetQuestionList(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// get question
func GetQuestion(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// submit code
func SubmitCode(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// get submissions
func GetSubmissions(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// get solution code
func GetSolutionCode(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// get leaderboard stats
func GetLeaderboardStats(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// get user stats
func GetUserStats(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// get user submissions
func GetUserSubmissions(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// get user solution code
func GetUserSolutionCode(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// get user leaderboard stats
func GetUserLeaderboardStats(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}
