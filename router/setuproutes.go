/***********************************************************************
     Copyright (c) 2025 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package router

import (
	"oj-backend/handler"
	"oj-backend/util"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// hello for root
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi there! This is the backend server for the Online Judge.")
	})

	api := app.Group("/api")

	// admin
	admin := api.Group("/admin")
	admin.Post("/login", handler.AdminLogin)
	admin.Use(util.AdminMiddleware())
	admin.Post("/create-admin", handler.CreateAdmin)
	admin.Put("/modify-languages", handler.ModifyLanguages)
	admin.Post("/create-contest", handler.CreateContest)
	admin.Post("/toggle-contest/:contestId", handler.ToggleContest)
	admin.Put("/modify-contest", handler.ModifyContest)
	admin.Post("/add-question", handler.AddQuestion)
	admin.Put("/modify-question", handler.ModifyQuestion)
	admin.Post("/add-testcase", handler.AddTestcase)
	admin.Delete("/delete-language", handler.DeleteLanguage)
	admin.Delete("/delete-contest/:contestId", handler.DeleteContest)
	admin.Delete("/delete-question/:questionId", handler.DeleteQuestion)
	admin.Delete("/delete-testcase/:testcaseId", handler.DeleteTestcase)
	admin.Delete("/delete-user/:userId", handler.DeleteUser)
	admin.Get("/get-languages", handler.GetLanguages)
	admin.Get("/get-contests", handler.GetContests)
	admin.Get("/get-questions/:contestId", handler.GetQuestions)
	admin.Get("/get-question-details", handler.GetQuestionDetails)
	admin.Get("/get-testcases/:questionId", handler.GetTestcases)
	admin.Get("/download-testcase-input/:testcaseId", handler.DownloadTestcaseInput)
	admin.Get("/download-testcase-output/:testcaseId", handler.DownloadTestcaseOutput)
	admin.Get("/get-userlist", handler.GetUserList)
	admin.Get("/get-submissionlist", handler.GetSubmissionList)
	admin.Get("/get-submission-code", handler.GetSubmissionCode)

	// auth
	auth := api.Group("/auth")
	auth.Post("/signup", handler.Signup)
	auth.Post("/login", handler.Login)
	auth.Get("/logout", handler.Logout)
	auth.Get("/reclaim-accesstoken", handler.ReclaimAccessToken)

	// api
	contest := api.Group("/contest")
	contest.Get("/get-contestlist", handler.GetContestList)
	contest.Get("/get-questionlist", handler.GetQuestionList)
	contest.Get("/get-question", handler.GetQuestion)
	contest.Post("/submit-code", handler.SubmitCode)
	contest.Get("/get-submissions", handler.GetSubmissions)
	contest.Get("/get-solutioncode", handler.GetSolutionCode)
	contest.Get("/get-leaderboardstats", handler.GetLeaderboardStats)
}
