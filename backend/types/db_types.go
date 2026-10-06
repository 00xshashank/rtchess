package types

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type GameModel struct {
	Player1   UserModel
	Player2   UserModel
	Moves     []string
	StartedAt time.Time
	WhiteWon  bool
}

func (UserModel) TableName() string {
	return "users"
}
