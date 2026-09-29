package auth

import (
	"errors"

	"github.com/00xshashank/rtchess/backend/db"
	"github.com/00xshashank/rtchess/backend/types"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 10)
}

func checkPasswordHash(password string, hash string) (bool, error) {
	pass_hash, pass_err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if pass_err != nil {
		return false, pass_err
	}

	return (string(pass_hash) == hash), nil
}

func signup(user types.User) error {
	hashed_password, pwd_err := hashPassword(user.Password)
	if pwd_err != nil {
		return pwd_err
	}

	user.Password = string(hashed_password)
	return db.AddUser(user)
}

func checkUserValidity(user types.User) error {
	if len(user.Username) < 6 {
		return errors.New("Username length should be at least 6")
	}

	if len(user.Password) < 8 {
		return errors.New("Password length must be at least 8")
	}
	if user.Firstname == "" {
		return errors.New("Invalid first name")
	}

	if user.Lastname == "" {
		return errors.New("Inalid last name")
	}

	if user.Email == "" {
		return errors.New("Invalid email")
	}

	return nil
}
