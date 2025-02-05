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

	api := app.Group("/api")

	// admin
	admin := api.Group("/admin")
	admin.Post("/login", handler.AdminLogin)
	admin.Use(util.AdminMiddleware())
	admin.Post("/create-admin", handler.CreateAdmin)
	admin.Put("/modify-languages", handler.ModifyLanguages)
	admin.Post("/create-contest", handler.CreateContest)
	admin.Put("/modify-contest", handler.ModifyContest)
	admin.Post("/add-question", handler.AddQuestion)
	admin.Put("/modify-question", handler.ModifyQuestion)
	admin.Post("/add-testcase", handler.AddTestcase)
	admin.Delete("/delete-language", handler.DeleteLanguage)
	admin.Delete("/delete-contest", handler.DeleteContest)
	admin.Delete("/delete-question", handler.DeleteQuestion)
	admin.Delete("/delete-testcase", handler.DeleteTestcase)
	admin.Get("/get-languages", handler.GetLanguages)
	admin.Get("/get-contests", handler.GetContests)
	admin.Get("/get-questions", handler.GetQuestions)
	admin.Get("/get-question-details", handler.GetQuestionDetails)
	admin.Get("/get-testcases", handler.GetTestcases)

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
