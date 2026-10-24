package db

import (
	"fmt"

	"github.com/00xshashank/rtchess/backend/backend/types"
)

func CreateGame(player1 types.User, player2 types.User) (uint, error) {
	player1db, player1db_err := GetUserFromDB(player1.Username)
	if player1db_err != nil {
		return 0, fmt.Errorf("Error while fetching details of player 1: %s\n", player1db_err)
	}

	player2db, player2db_err := GetUserFromDB(player2.Username)
	if player2db_err != nil {
		return 0, fmt.Errorf("Error while fetching details of player 2: %s\n", player2db_err)
	}

	game := types.GameModel{
		Player1ID:         player1db.ID,
		Player2ID:         player2db.ID,
		Player1PlaysWhite: true,
		Result:            uint(types.ResultGameInProgress),
	}

	result := db.Create(game)
	if result.Error != nil {
		return 0, result.Error
	}

	return game.ID, nil
}
