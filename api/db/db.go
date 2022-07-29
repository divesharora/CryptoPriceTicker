package db

import (
	// "database/sql"
	"fmt"
	"log"
	// "time"

	"github.com/divesharora/KryptoBackendTask/api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/lib/pq"
	"github.com/spf13/viper"
	// "gorm.io/driver/postgres"
)

var db *gorm.DB = nil

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}

	db = Connect()
	return db
}

func Connect() *gorm.DB {
	username := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASS")
	dbName := viper.GetString("DB_NAME")
	dbHost := viper.GetString("DB_HOST")
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s port=5432 password=%s sslmode=disable", dbHost, username, dbName, password)
	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Users{})
	return db
}
