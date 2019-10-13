package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

func DatabaseHandler() *gorm.DB {
	databaseUrl := os.Getenv("DATABASE_URL")
	if len(databaseUrl) == 0 {
		databaseUrl = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			os.Getenv("POSTGRESQL_HOST"),
			os.Getenv("POSTGRESQL_PORT"),
			os.Getenv("POSTGRESQL_USER"),
			os.Getenv("POSTGRESQL_DBNAME"),
			os.Getenv("POSTGRESQL_PASSWORD"),
		)
	}

	db, err := gorm.Open("postgres", databaseUrl)

	if err != nil {
		fmt.Println(err)
		panic("Error establishing connection to DB.")
	}

	return db
}