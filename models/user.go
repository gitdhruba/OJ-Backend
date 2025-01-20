/***********************************************************************
     Copyright (c) 2024 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package models

import "github.com/golang-jwt/jwt/v5"

// user model
type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

type Admin User

// jwt claims structure
type Claims struct {
	ClaimID        int `json:"id" gorm:"primaryKey"`
	StandardClaims jwt.RegisteredClaims
	Admin          bool `gorm:"not null"`
}

// GetAudience implements jwt.Claims.
func (c Claims) GetAudience() (jwt.ClaimStrings, error) {
	return c.StandardClaims.GetAudience()
}

// GetExpirationTime implements jwt.Claims.
func (c Claims) GetExpirationTime() (*jwt.NumericDate, error) {
	return c.StandardClaims.GetExpirationTime()
}

// GetIssuedAt implements jwt.Claims.
func (c Claims) GetIssuedAt() (*jwt.NumericDate, error) {
	return c.StandardClaims.GetIssuedAt()
}

// GetIssuer implements jwt.Claims.
func (c Claims) GetIssuer() (string, error) {
	return c.StandardClaims.GetIssuer()
}

// GetNotBefore implements jwt.Claims.
func (c Claims) GetNotBefore() (*jwt.NumericDate, error) {
	return c.StandardClaims.GetNotBefore()
}

// GetSubject implements jwt.Claims.
func (c Claims) GetSubject() (string, error) {
	return c.StandardClaims.GetSubject()
}

// dummy Validate method
func (c Claims) Validate() error {
	return nil
}
