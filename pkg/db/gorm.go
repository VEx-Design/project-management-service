package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresdb struct {
	DB *gorm.DB
}

func ConnectToPG() *postgresdb {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s ",
		os.Getenv("PROJECT_DB_HOST"),
		os.Getenv("PROJECT_DB_USERNAME"),
		os.Getenv("PROJECT_DB_PASSWORD"),
		os.Getenv("PROJECT_DB_NAME"),
		os.Getenv("PROJECT_DB_PORT"),
	)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Fail to connect to DB")
	}
	return &postgresdb{DB: DB}
}

func (db *postgresdb) GetClient() *gorm.DB {
	return db.DB
}
