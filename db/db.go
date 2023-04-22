package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	RecordNotFound = "record not found"
)

var models = []any{
	&Post{},
	&Comment{},
	&User{},
}

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tofus-blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(models...)
}

func GetDB() *gorm.DB {
	return DB
}
