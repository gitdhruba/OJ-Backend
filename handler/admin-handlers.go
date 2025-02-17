/***********************************************************************
     Copyright (c) 2025 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package handler

import (
	"fmt"
	"oj-backend/config"
	db "oj-backend/database"
	"oj-backend/models"
	"oj-backend/util"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// login
func AdminLogin(c *fiber.Ctx) error {
	var u models.Admin
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	password := u.Password
	// fmt.Println(db.DB)
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
	var u models.Admin
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if u.Username == "" || u.Email == "" || u.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "username, email and password are required",
		})
	}

	if res := db.DB.Where(&models.Admin{Username: u.Username}).Or(&models.Admin{Email: u.Email}).First(&u); (res.Error == nil) || (res.RowsAffected > 0) {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "username or email already exists",
		})
	}

	cost, err := strconv.Atoi(config.GetEnv("BCRYPT_COST"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not create admin",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), cost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not create admin",
		})
	}

	u.Password = string(hashedPassword)
	if res := db.DB.Model(&u).Create(&u); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not create admin",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "admin created successfully",
	})
}

// create-contest
func CreateContest(c *fiber.Ctx) error {
	var contest models.Contest
	if err := c.BodyParser(&contest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	// if contest.Title == "" || len(contest.Description) == 0 || contest.StartTime == "" || contest.EndTime == "" || len(contest.Languages) == 0 {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": "title, description, start_time, end_time and languages are required",
	// 	})
	// }

	// if res := db.DB.Where(&models.Contest{Title: contest.Title}).First(&contest); (res.Error == nil) || (res.RowsAffected > 0) {
	// 	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
	// 		"error": "contest already exists",
	// 	})
	// }

	if res := db.DB.Save(&contest); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": res.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "contest created successfully",
		"contestID": contest.ID,
	})
}

// toggle-contest
func ToggleContest(c *fiber.Ctx) error {
	contestID, err := strconv.Atoi(c.Params("contestId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid contest_id",
		})
	}

	var contest models.Contest
	if res := db.DB.Where(&models.Contest{ID: contestID}).First(&contest); (res.Error != nil) || (res.RowsAffected <= 0) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "contest does not exist",
		})
	}

	contest.IsActive = !contest.IsActive
	if res := db.DB.Model(&contest).Updates(&models.Contest{IsActive: contest.IsActive}); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not toggle contest mode",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// modify-contest
func ModifyContest(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// update-languages
func ModifyLanguages(c *fiber.Ctx) error {
	fmt.Println("ModifyLanguages")
	return c.SendStatus(fiber.StatusForbidden)
}

// add-question
func AddQuestion(c *fiber.Ctx) error {
	var question models.Question
	if err := c.BodyParser(&question); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if res := db.DB.Save(&question); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": res.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "question created successfully",
		"questionID": question.ID,
	})
}

// update-question
func ModifyQuestion(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// add-testcase
func AddTestcase(c *fiber.Ctx) error {
	inputFile, err1 := c.FormFile("input_file")
	outputFile, err2 := c.FormFile("output_file")

	if err1 != nil || err2 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Both input.txt and output.txt must be uploaded",
		})
	}

	contestID, err3 := strconv.Atoi(c.FormValue("contest_id", "NA"))
	questionID, err4 := strconv.Atoi(c.FormValue("question_id", "NA"))
	testcaseNo, err5 := strconv.Atoi(c.FormValue("testcase_no", "NA"))

	if err3 != nil || err4 != nil || err5 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid contest_id or question_id",
		})
	}

	// check whether contest and question exists
	var contest models.Contest
	if res := db.DB.Where(&models.Contest{ID: contestID}).First(&contest); (res.Error != nil) || (res.RowsAffected <= 0) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "contest does not exist",
		})
	}

	var question models.Question
	if res := db.DB.Where(&models.Question{ID: questionID}).First(&question); (res.Error != nil) || (res.RowsAffected <= 0) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "question does not exist",
		})
	}

	// create testcase directory
	testcasePath := fmt.Sprintf("%s/%d/%d", config.GetEnv("CONTEST_PATH"), contestID, questionID)
	if err := os.MkdirAll(testcasePath, 0755); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not create testcase directory",
		})
	}

	inputFile.Filename = fmt.Sprintf("input_%d.txt", testcaseNo)
	outputFile.Filename = fmt.Sprintf("output_%d.txt", testcaseNo)
	testcase := models.Testcase{
		No:             testcaseNo,
		QuestionID:     questionID,
		InputFilePath:  fmt.Sprintf("%s/%s", testcasePath, inputFile.Filename),
		OutputFilePath: fmt.Sprintf("%s/%s", testcasePath, outputFile.Filename),
	}

	// save input file
	if err := c.SaveFile(inputFile, testcase.InputFilePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not save input file",
		})
	}

	// save output file
	if err := c.SaveFile(outputFile, testcase.OutputFilePath); err != nil {
		// delete input file
		os.Remove(fmt.Sprintf("%s/%s", testcasePath, inputFile.Filename))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not save output file",
		})
	}

	// save testcase in database
	if res := db.DB.Save(&testcase); res.Error != nil {
		// delete input and output files
		os.Remove(testcase.InputFilePath)
		os.Remove(testcase.OutputFilePath)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": res.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "testcase created successfully",
		"testcaseID": testcase.ID,
	})
}

// delete-language
func DeleteLanguage(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// delete-contest
func DeleteContest(c *fiber.Ctx) error {
	contestId, err := strconv.Atoi(c.Params("contestId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid contest_id",
		})
	}

	if res := db.DB.Delete(&models.Contest{ID: contestId}); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": res.Error.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// delete-question
func DeleteQuestion(c *fiber.Ctx) error {
	questionId, err := strconv.Atoi(c.Params("questionId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid question_id",
		})
	}

	if res := db.DB.Delete(&models.Question{ID: questionId}); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": res.Error.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// delete-testcase
func DeleteTestcase(c *fiber.Ctx) error {
	testcaseId, err := strconv.Atoi(c.Params("testcaseId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid testcase_id",
		})
	}

	if res := db.DB.Delete(&models.Testcase{ID: testcaseId}); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": res.Error.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// get-languages
func GetLanguages(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// get-contests
func GetContests(c *fiber.Ctx) error {
	var contests []models.Contest
	if res := db.DB.Find(&contests); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": res.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"contests": contests,
	})
}

// get-questions
func GetQuestions(c *fiber.Ctx) error {
	contestId, err := strconv.Atoi(c.Params("contestId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid contest_id",
		})
	}

	var questions []models.Question
	if res := db.DB.Where(&models.Question{ContestID: contestId}).Find(&questions); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": res.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"questions": questions,
	})
}

// get-question-details
func GetQuestionDetails(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// get-testcases
func GetTestcases(c *fiber.Ctx) error {
	questionId, err := strconv.Atoi(c.Params("questionId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid question_id",
		})
	}

	var testcases []models.Testcase
	if res := db.DB.Where(&models.Testcase{QuestionID: questionId}).Find(&testcases); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": res.Error.Error(),
		})
	}

	// set input and output file paths as input_no.txt and output_no.txt
	for i := range len(testcases) {
		testcases[i].InputFilePath = fmt.Sprintf("input_%d.txt", testcases[i].No)
		testcases[i].OutputFilePath = fmt.Sprintf("output_%d.txt", testcases[i].No)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"testcases": testcases,
	})
}
