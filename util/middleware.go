/***********************************************************************
     Copyright (c) 2024 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package util

import (
	"oj-backend/config"
	"oj-backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func validateToken(tokenstring string) (*models.Claims, bool) {
	jwtSecret := []byte(config.GetEnv("JWT_SECRET"))
	token, err := jwt.ParseWithClaims(tokenstring, &models.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

	if err != nil || !token.Valid {
		return nil, false
	}

	claims, ok := token.Claims.(*models.Claims)
	if !ok {
		return nil, false
	}

	return claims, true
}

// middleware for /api/
func ApiMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		accessToken := c.Get("access_token", "")
		refreshToken := c.Get("refresh_token", "")

		if accessToken == "" || refreshToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		accessClaims, ok := validateToken(accessToken)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		user, _ := accessClaims.GetIssuer()

		if expiry, _ := accessClaims.GetExpirationTime(); expiry.Time.Before(time.Now()) {
			refreshClaims, ok := validateToken(refreshToken)
			if !ok {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Unauthorized",
				})
			}

			if tmpuser, _ := refreshClaims.GetIssuer(); user != tmpuser {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Unauthorized",
				})
			}

			if expiry, _ := refreshClaims.GetExpirationTime(); expiry.Time.Before(time.Now()) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Unauthorized",
				})
			}

			// update access and refresh tokens
			accessToken, refreshToken, ok = GenerateTokens(user, false)
			if !ok {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Couldn't regenerate tokens",
				})
			}
		}

		// set access and refresh tokens in response header
		c.Set("access_token", accessToken)
		c.Set("refresh_token", refreshToken)

		// save userid for request processing
		c.Locals("user", user)

		return c.Next()
	}
}

// middleware for /admin/
func AdminMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		accessToken := c.Get("access_token", "")
		refreshToken := c.Get("refresh_token", "")

		if accessToken == "" || refreshToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		accessClaims, ok := validateToken(accessToken)
		if !ok || !accessClaims.Admin {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		user, _ := accessClaims.GetIssuer()

		if expiry, _ := accessClaims.GetExpirationTime(); expiry.Time.Before(time.Now()) {
			refreshClaims, ok := validateToken(refreshToken)
			if !ok || !refreshClaims.Admin {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Unauthorized",
				})
			}

			if tmpuser, _ := refreshClaims.GetIssuer(); user != tmpuser {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Unauthorized",
				})
			}

			if expiry, _ := refreshClaims.GetExpirationTime(); expiry.Time.Before(time.Now()) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Unauthorized",
				})
			}

			// update access and refresh tokens
			accessToken, refreshToken, ok = GenerateTokens(user, true)
			if !ok {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Couldn't regenerate tokens",
				})
			}
		}

		// set access and refresh tokens in response header
		c.Set("access_token", accessToken)
		c.Set("refresh_token", refreshToken)

		// save userid for request processing
		c.Locals("user", user)

		return c.Next()
	}
}
