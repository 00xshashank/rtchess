package types

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username  string `gorm:"unique" json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Rating    uint   `json:"rating"`
}

func (UserModel) TableName() string {
	return "users"
}

type GameModel struct {
	gorm.Model
	Player1ID         uint
	Player2ID         uint
	Player1PlaysWhite bool
	Result            uint

	Player1 UserModel `gorm:"foreignkey:Player1ID"`
	Player2 UserModel `gorm:"foreignkey:Player2ID"`
}

func (GameModel) TableName() string {
	return "games"
}

type MoveModel struct {
	GameID uint
	Game   GameModel `gorm:"foreignkey:GameID"`

	Index uint16
	Move  string
}

func (MoveModel) TableName() string {
	return "moves"
}
