package auth

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/00xshashank/rtchess/backend/backend/types"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type jwtClaims struct {
	jwt.RegisteredClaims
	Username  string `json:"usr"`
	Email     string `json:"eml"`
	Firstname string `json:"fnm"`
	Lastname  string `json:"lnm"`
	Rating    uint   `json:"rat"`
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 10)
}

func CheckPasswordHash(password []byte, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)

	if err != nil {
		return false
	}

	return true
}

func CheckUserValidity(user types.UserModel) error {
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

func CreateJWT(user types.UserModel) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "bsdk",
		"usr": user.Username,
		"eml": user.Email,
		"fnm": user.Firstname,
		"lnm": user.Lastname,
		"rat": user.Rating,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	key, err := token.SignedString([]byte(secret))

	return key, err
}

func GetUserFromCookie(cookie *http.Cookie) (types.User, error) {
	token, token_err := jwt.ParseWithClaims(cookie.Value, &jwtClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if token_err != nil {
		return types.User{}, token_err
	}

	exp_time, exp_err := token.Claims.GetExpirationTime()
	if exp_err != nil {
		log.Printf("Error while getting expiration time: %s\n", exp_err)
		return types.User{}, exp_err
	}

	if exp_time.Time.Unix() < time.Now().Unix() {
		return types.User{}, errors.New("Session Expired")
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return types.User{}, errors.New("Failed to parse cookie")
	}

	var user types.User = types.User{
		Firstname: claims.Username,
		Lastname:  claims.Lastname,
		Email:     claims.Email,
		Username:  claims.Username,
		Rating:    claims.Rating,
	}

	return user, nil
}
