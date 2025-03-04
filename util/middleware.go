/***********************************************************************
     Copyright (c) 2025 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package util

import (
	"oj-backend/config"
	"oj-backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func validateToken(tokenstring string) (*models.Claims, bool) {
	jwtSecret := []byte(config.GetEnv("JWT_SECRET"))
	token, err := jwt.ParseWithClaims(tokenstring, &models.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

	claims, ok := token.Claims.(*models.Claims)
	if !ok {
		return nil, false
	}

	return claims, ((err == nil) && token.Valid)
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

		accessClaims, isValid := validateToken(accessToken)
		if accessClaims == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		user, _ := accessClaims.GetIssuer()

		if !isValid {
			refreshClaims, ok := validateToken(refreshToken)
			if (refreshClaims == nil) || !ok {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Unauthorized",
				})
			}

			if tmpuser, _ := refreshClaims.GetIssuer(); user != tmpuser {
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

			// set new access and refresh tokens in response header
			c.Set("access_token", accessToken)
			c.Set("refresh_token", refreshToken)
		}

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

		accessClaims, isValid := validateToken(accessToken)
		if (accessClaims == nil) || !accessClaims.Admin {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		user, _ := accessClaims.GetIssuer()

		if !isValid {
			refreshClaims, ok := validateToken(refreshToken)
			if (refreshClaims == nil) || !ok || !refreshClaims.Admin {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Unauthorized",
				})
			}

			if tmpuser, _ := refreshClaims.GetIssuer(); user != tmpuser {
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

			// set access and refresh tokens in response header
			c.Set("access_token", accessToken)
			c.Set("refresh_token", refreshToken)
		}

		// save userid for request processing
		c.Locals("user", user)

		return c.Next()
	}
}
