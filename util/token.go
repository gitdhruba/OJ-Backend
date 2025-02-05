/***********************************************************************
     Copyright (c) 2025 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package util

import (
	"oj-backend/config"
	"oj-backend/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateTokens returns the access and refresh tokens for a given user ID and isAdmin status
func GenerateTokens(id string, isAdmin bool) (string, string, bool) {
	t := time.Now()
	jwtSecret := []byte(config.GetEnv("JWT_SECRET"))
	jwtClaim := models.Claims{
		StandardClaims: jwt.RegisteredClaims{
			Issuer:    id,
			ExpiresAt: jwt.NewNumericDate(t.Add(time.Minute * 15)),
			Subject:   "access_token",
			IssuedAt:  jwt.NewNumericDate(t),
		},
		Admin: isAdmin,
	}

	// generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", false
	}

	// generate refresh token
	jwtClaim.StandardClaims.Subject = "refresh_token"
	jwtClaim.StandardClaims.ExpiresAt = jwt.NewNumericDate(t.Add(time.Hour * 24))
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	refreshTokenString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", false
	}

	return accessTokenString, refreshTokenString, true
}
