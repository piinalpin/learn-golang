package config

import (
	"log"
	"os"
	"time"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	var db = connectDB()
	return db
}

func connectDB() *gorm.DB {
	godotenv.Load()
	var err error
	var dsn = os.ExpandEnv("${datasource.username}:${datasource.password}@tcp(${datasource.url})/${datasource.name}?parseTime=true&loc=Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
		return nil
	}

	sqlDB, _ := db.DB()
	mid, _ := strconv.Atoi(os.Getenv("pool.set-max-idle-connection"))
	mao, _ := strconv.Atoi(os.Getenv("pool.set-max-open-connection"))
	mlc, _ := strconv.Atoi(os.Getenv("pool.set-max-lifetime-connection"))
	sqlDB.SetMaxIdleConns(mid)
	sqlDB.SetMaxOpenConns(mao)
	sqlDB.SetConnMaxLifetime(time.Duration(mlc) * 60000 * time.Millisecond)

	return db
}
