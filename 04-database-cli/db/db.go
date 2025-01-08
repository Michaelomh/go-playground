package db

import (
	"log"

	_ "github.com/glebarez/sqlite"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Driver   string
	Name     string
	Host     string
	User     string
	Password string
}

var dbInstance *gorm.DB

func Get() *gorm.DB {
	return dbInstance
}

func init() {
	// _ := Config{
	// 	Driver:   os.Getenv("DB_DRIVER"),
	// 	Name:     os.Getenv("DB_NAME"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	Host:     os.Getenv("DB_HOST"),
	// 	User:     os.Getenv("DB_USER"),
	// }

	var err error
	dbInstance, err = gorm.Open(sqlite.Open("app_db.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
}
