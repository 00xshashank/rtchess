package types

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Game struct {
	Player1   User
	Player2   User
	Moves     []string
	StartedAt time.Time
	WhiteWon  bool
}
