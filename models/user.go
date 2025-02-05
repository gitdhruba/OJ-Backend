/***********************************************************************
     Copyright (c) 2025 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package models

// user model
type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

// use same model for admin
type Admin User
