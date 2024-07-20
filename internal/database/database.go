package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Service struct {
	Db *gorm.DB
}

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	url        = os.Getenv("URL")
	dbInstance *Service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return *dbInstance
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, url, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	dbInstance = &Service{
		Db: db,
	}
	return *dbInstance
}
