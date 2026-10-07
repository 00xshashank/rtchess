package db

import (
	"log"

	"github.com/00xshashank/rtchess/backend/types"
)

func CreateGame(player1 types.User, player2 types.User) (uint, error) {
	// player1id :=

	game := types.GameModel{
		// Player1ID uint
		// Player1   UserModel `gorm:"foreignkey:Player1ID"`

		// Player2ID uint
		// Player2   UserModel `gorm:"foreignkey:Player2ID"`

		// Result uint
	}
	log.Printf("%v", game)

	return 0, nil
}
