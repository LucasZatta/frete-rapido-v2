package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type Service struct {
// 	Db *gorm.DB
// }

// var (
// 	database   = os.Getenv("DB_DATABASE")
// 	password   = os.Getenv("DB_PASSWORD")
// 	username   = os.Getenv("DB_USERNAME")
// 	port       = os.Getenv("DB_PORT")
// 	host       = os.Getenv("DB_HOST")
// 	schema     = os.Getenv("DB_SCHEMA")
// 	dbInstance *Service
// )

// func New() Service {
// 	if dbInstance != nil {
// 		fmt.Println("here")
// 		return *dbInstance
// 	}

// 	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, username, password, database, port, "disable", "Asia/Shanghai")
// 	// connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)
// 	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

// 	// sqlDB, err := sql.Open("pgx", connStr)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// db, err := gorm.Open(postgres.New(postgres.Config{
// 	// 	Conn: sqlDB,
// 	// }), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	dbInstance = &Service{
// 		Db: db,
// 	}
// 	return *dbInstance
// }

type service struct {
	db *gorm.DB
}

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	schema     = os.Getenv("DB_SCHEMA")
	dbInstance *service
)

func New() *gorm.DB {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance.db
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	dbInstance = &service{
		db: gormDB,
	}
	return gormDB
}
