package config

import (
	"log"
	"os"

	"learn-rest-api/cmd/app/domain/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)

var Db *gorm.DB

func InitDB() *gorm.DB {
	var err error
	var dsn = os.ExpandEnv("${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?parseTime=true&loc=Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
		return nil
	}

	db.AutoMigrate(&dao.Author{})

	return db
}
