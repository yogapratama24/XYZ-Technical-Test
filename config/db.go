package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"xyz_test/model"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")

	var err error

	migrate, err := strconv.ParseBool(os.Getenv("MIGRATE"))
	if err != nil {
		panic("Seeder status must be filled")
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost, dbPort, username, password, dbname)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect database")
	}

	sqlDB, _ := DB.DB()
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Database not responding with err: %s", err)
	}

	if migrate {
		DB.AutoMigrate(
			&model.User{},
			&model.Transaction{},
		)
	}
	log.Println("Success connected database")
	return DB
}
