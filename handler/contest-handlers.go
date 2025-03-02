/***********************************************************************
     Copyright (c) 2025 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package handler

import (
	db "oj-backend/database"
	"oj-backend/models"

	"github.com/gofiber/fiber/v2"
)

// get contest-list
func GetContestList(c *fiber.Ctx) error {
	var contests []models.Contest
	if err := db.DB.Find(&contests).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching contests",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"contests": contests,
	})
}

// get question-list
func GetQuestionList(c *fiber.Ctx) error {
	contestId, err := c.ParamsInt("contestId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid contest id",
		})
	}

	var questions []models.Question
	// fields := []string{"id", "name", "points"}
	if err := db.DB.Where(&models.Question{ContestID: contestId}).Order("points").Find(&questions).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	submissionCnt := make([]int64, len(questions))
	submissionCntAC := make([]int64, len(questions))
	for i, question := range questions {
		db.DB.Model(&models.Submission{}).Where(&models.Submission{QuestionID: question.ID}).Count(&submissionCnt[i])
		db.DB.Model(&models.Submission{}).Where(&models.Submission{QuestionID: question.ID, Verdict: "AC"}).Count(&submissionCntAC[i])
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"questions":           questions,
		"submission_count":    submissionCnt,
		"submission_count_ac": submissionCntAC,
	})
}

// get question
func GetQuestion(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// submit code
func SubmitCode(c *fiber.Ctx) error {
	var submission models.Submission
	if err := c.BodyParser(&submission); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid submission details",
		})
	}

	/*
		send the code to task queue
	*/

	if err := db.DB.Create(&submission).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error submitting code",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Code submitted successfully",
		"submissionId": submission.ID,
	})
}

// get submissions
func GetSubmissions(c *fiber.Ctx) error {
	var req struct {
		QuestionId int `json:"question_id"`
		UserId     int `json:"user_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	var submissions []models.Submission
	fields := []string{"id", "submission_time", "verdict", "exec_time", "memory"}
	if err := db.DB.Select(fields).Where(&models.Submission{
		QuestionID: req.QuestionId,
		UserID:     req.UserId,
	}).Find(&submissions).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching submissions",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"submissions": submissions,
	})
}

// get solution code
func GetSolutionCode(c *fiber.Ctx) error {
	submissionId, err := c.ParamsInt("submissionId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid submission id",
		})
	}

	var submission models.Submission
	fields := []string{"code"}
	if err := db.DB.Select(fields).Where(&models.Submission{ID: submissionId}).First(&submission).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching solution code",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": string(submission.Code),
	})
}

// get leaderboard stats
func GetLeaderboardStats(c *fiber.Ctx) error {
	contestId, err := c.ParamsInt("contestId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid contest id",
		})
	}

	// get question ids for the contest oredered by points
	var questionIds []int
	if err := db.DB.Model(&models.Question{}).Where(&models.Question{ContestID: contestId}).Order("points").Pluck("id", &questionIds).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching questions",
		})
	}

	var leaderboard []models.Leaderboard

	// get users according to points(inc) ans penalty(dec) from scores join users
	rows, err := db.DB.Table("scores").Select("users.id users.username, scores.points, scores.penalty").Joins("join users on users.id = scores.user_id").Where("contest_id = ?", contestId).Order("points desc, penalty asc").Rows()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching leaderboard",
		})
	}

	defer rows.Close()
	for rows.Next() {
		var user_id int
		var username string
		var points, penalty int
		rows.Scan(&user_id, &username, &points, &penalty)

		// get submission count for each question
		// make submission count negative if there is no AC verdict
		submissionCnt := make([]int64, len(questionIds))
		fields := []string{"id", "verdict"}
		for i, questionId := range questionIds {
			// total submissions
			db.DB.Model(&models.Submission{}).Select(fields).Where(&models.Submission{QuestionID: questionId, UserID: user_id}).Count(&submissionCnt[i])
			// AC submissions
			var acCnt int64
			db.DB.Model(&models.Submission{}).Select(fields).Where(&models.Submission{QuestionID: questionId, UserID: user_id, Verdict: "AC"}).Count(&acCnt)
			if acCnt == 0 {
				submissionCnt[i] = -submissionCnt[i]
			}
		}

		leaderboard = append(leaderboard, models.Leaderboard{
			UserID:   user_id,
			Username: username,
			Points:   points,
			Penalty:  penalty,
			Trials:   submissionCnt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"leaderboard": leaderboard,
	})
}

// get user stats
func GetUserStats(c *fiber.Ctx) error {
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
