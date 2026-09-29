package db

import (
	"fmt"
	"os"

	"github.com/00xshashank/rtchess/backend/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	// fmt.Printf("dsn: %s\n", dsn)

	var db_err error
	db, db_err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if db_err != nil {
		fmt.Printf("Error while opening connection to a database: %s\n", db_err)
		return db_err
	}

	return nil
}

func AddUser(user types.User) error {
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUser(username string) (types.User, error) {
	var user types.User

	result := db.First(&user, "username = ?", username)
	if result.Error != nil {
		return types.User{}, result.Error
	}

	return user, nil
}
