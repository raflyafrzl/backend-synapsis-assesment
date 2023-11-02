package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

func NewInitDB(dsn string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(os.Getenv("DATA_SOURCE_NAME")), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	sql, err := db.DB()

	sql.SetConnMaxIdleTime(time.Minute * 5)
	sql.SetConnMaxLifetime(time.Hour)
	sql.SetMaxIdleConns(3)
	sql.SetMaxOpenConns(10)

	return db
}
