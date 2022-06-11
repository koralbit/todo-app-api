package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgreSQL(config PostgreSQLConfiguration) *gorm.DB {
	if db != nil {
		return db
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Password, config.Host, config.Port, config.Database)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	return db
}

type PostgreSQLConfiguration struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}
