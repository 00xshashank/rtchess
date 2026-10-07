package db

import (
	"fmt"
	"os"

	"github.com/00xshashank/rtchess/backend/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	var db_err error
	db, db_err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if db_err != nil {
		fmt.Printf("Error while opening connection to a database: %s\n", db_err)
		return db_err
	}

	db.AutoMigrate(
		&types.UserModel{},
		&types.GameModel{},
		&types.MoveModel{},
	)

	return nil
}
