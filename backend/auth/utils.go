package auth

import (
	"errors"

	"github.com/00xshashank/rtchess/backend/types"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 10)
}

func checkPasswordHash(password []byte, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)

	if err != nil {
		return false
	}

	return true
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
