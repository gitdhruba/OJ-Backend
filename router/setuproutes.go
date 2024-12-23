/***********************************************************************
     Copyright (c) 2024 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package router

import (
	"oj-backend/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// admin
	admin := app.Group("/admin")
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
	auth := app.Group("/auth")

	// api
	api := app.Group("/api")

}
