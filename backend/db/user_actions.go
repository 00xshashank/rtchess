package db

import "github.com/00xshashank/rtchess/backend/types"

func AddUserToDB(user types.UserModel) error {
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUserFromDB(username string) (types.UserModel, error) {
	var user types.UserModel

	result := db.First(&user, "username = ?", username)
	if result.Error != nil {
		return types.UserModel{}, result.Error
	}

	return user, nil
}
