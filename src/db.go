package src

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

var models = []any{
	&Post{},
	&Comment{},
	&User{},
	&Permission{},
}

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("tofus-blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(models...)
}

func GetDB() *gorm.DB {
	return db
}
