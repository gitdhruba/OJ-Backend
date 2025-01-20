/***********************************************************************
     Copyright (c) 2024 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// path for env file
const configPath = "./.env"

// Load environment variables from env file
func LoadEnv() {
	fmt.Println("[-->] loading env")
	err := godotenv.Load(configPath)
	if err != nil {
		panic(err)
	}
}

// retrieve env value from key
func GetEnv(key string) string {
	val, isSet := os.LookupEnv(key)
	if !isSet {
		panic(fmt.Sprintf("[ERROR] environment variable %s is not set\n", key))
	}

	return val
}
