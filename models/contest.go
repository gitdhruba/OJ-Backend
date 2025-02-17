/***********************************************************************
     Copyright (c) 2025 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package models

import (
	"errors"
	"os"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

// contest model
type Contest struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	Title       string         `json:"name" gorm:"unique, not null"`
	Description []byte         `json:"description" gorm:"type:bytea"`
	StartTime   string         `json:"start_time" gorm:"not null"`
	EndTime     string         `json:"end_time" gorm:"not null"`
	Languages   pq.StringArray `json:"languages" gorm:"type:text[]"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
}

func (contest *Contest) BeforeSave(tx *gorm.DB) error {
	startTime, err1 := time.Parse(time.RFC3339, contest.StartTime)
	endTime, err2 := time.Parse(time.RFC3339, contest.EndTime)
	if contest.Title == "" || len(contest.Description) == 0 || len(contest.Languages) == 0 || err1 != nil || err2 != nil || startTime.After(endTime) {
		return errors.New("invalid contest details")
	}

	return nil
}

// Question model
type Question struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	ContestID   int     `json:"contest_id"`
	Contest     Contest `gorm:"foreignKey:ContestID, references:ID, constraint:OnUpdate:CASCADE, OnDelete:CASCADE"`
	Name        string  `json:"name" gorm:"not null"`
	Description []byte  `json:"description" gorm:"not null, type:bytea"`
	Points      int     `json:"points" gorm:"not null"`
	TimeLimit   int     `json:"time_limit" gorm:"not null"`
	MemoryLimit int     `json:"memory_limit" gorm:"not null"`
}

func (question *Question) BeforeSave(tx *gorm.DB) error {
	if question.Points <= 0 || question.TimeLimit <= 0 || question.MemoryLimit <= 0 {
		return errors.New("invalid points, time_limit or memory_limit")
	}

	return nil
}

// testcase model
type Testcase struct {
	ID             int      `json:"id" gorm:"primaryKey"`
	No             int      `json:"no" gorm:"not null"`
	QuestionID     int      `json:"question_id"`
	Question       Question `gorm:"foreignKey:QuestionID, references:ID, constraint:OnUpdate:CASCADE, OnDelete:CASCADE"`
	InputFilePath  string   `json:"input_file_path" gorm:"not null"`
	OutputFilePath string   `json:"output_file_path" gorm:"not null"`
}

func (testcase *Testcase) BeforeDelete(tx *gorm.DB) error {
	os.Remove(testcase.InputFilePath)
	os.Remove(testcase.OutputFilePath)
	return nil
}

// submission model
type Submission struct {
	ID             int      `json:"id" gorm:"primaryKey"`
	QuestionID     int      `json:"question_id"`
	Question       Question `gorm:"foreignKey:QuestionID, references:ID, constraint:OnUpdate:CASCADE, OnDelete:CASCADE"`
	UserID         int      `json:"user_id"`
	User           User     `gorm:"foreignKey:UserID, references:ID, constraint:OnUpdate:CASCADE, OnDelete:CASCADE"`
	Language       string   `json:"language" gorm:"not null"`
	Code           []byte   `json:"code" gorm:"not null, type:bytea"`
	SubmissionTime string   `json:"submission_time" gorm:"not null"`
	Verdict        string   `json:"verdict" gorm:"not null"`
}
