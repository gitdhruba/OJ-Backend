/***********************************************************************
     Copyright (c) 2025 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package util

import (
	"oj-backend/config"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// genearte hash of password
func HashPassword(password *string) bool {
	cost, err := strconv.Atoi(config.GetEnv("BCRYPT_COST"))
	if err != nil {
		return false
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), cost)
	if err != nil {
		return false
	}

	*password = string(hashedPassword)
	return true
}

// compare password with hash
func ComparePassword(hashedPassword, password string) bool {
	return (bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil)
}
