/***********************************************************************
     Copyright (c) 2025 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package models

import "github.com/golang-jwt/jwt/v5"

// jwt claims structure
type Claims struct {
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
