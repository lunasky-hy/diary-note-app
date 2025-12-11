package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lunasky-hy/dialy-note-app/src/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() *gorm.DB {
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_PASS := os.Getenv("DB_PASSWORD")
	if DB_PASS == "" {
		DB_PASS="null"
	}
	DB_URL := os.Getenv("DB_URL")
	DB_PORTNUM, _ := strconv.Atoi(DB_PORT)

	dsn := DB_URL
	if DB_URL == ""{
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", 
			DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORTNUM)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if (err != nil) {
		fmt.Println(err.Error())
	}

	db.AutoMigrate(&model.Question{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Diary{})
	return db
}