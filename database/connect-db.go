/***********************************************************************
     Copyright (c) 2024 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package database

import (
	"fmt"
	"oj-backend/config"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// global DB variable for accessing database
var DB *gorm.DB

// connect to DB instance
func ConnectDB() {
	dbHost := config.GetEnv("DB_HOST")
	dbPort, err := strconv.ParseUint(config.GetEnv("DB_PORT"), 10, 32)
	if err != nil {
		panic("[ERROR] invalid DB-port")
	}

	dbName := config.GetEnv("DB_NAME")
	dbUser := config.GetEnv("DB_USER")
	dbPassword := config.GetEnv("DB_PASSWORD")
	// dbSSLMode := config.GetEnv("DB_SSL")

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPassword, dbName)

	DB, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("[-->] connected to database")

	DB.Logger = logger.Default.LogMode(logger.Info)

	// automigration
	// pending
}
