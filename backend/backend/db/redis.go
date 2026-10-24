package db

import (
	"context"
	"log"
	"math"
	"strconv"

	"github.com/00xshashank/rtchess/backend/backend/types"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func checkForOpponent(user types.User) (string, error) {
	log.Printf("Checking for opponents for: %s\n", user.Username)

	result, err := RedisClient.ZRangeByScoreWithScores(
		context.Background(),
		"game",
		&redis.ZRangeBy{
			Min: strconv.FormatInt(int64(math.Max(float64(user.Rating-100), 0)), 10),
			Max: strconv.FormatInt(int64(user.Rating+100), 10),
		},
	).Result()

	log.Printf("Obtained result from zrange: %v\n", result)

	if err != nil {
		return "", err
	}

	if len(result) < 1 {
		return "", nil
	}

	result_uname := result[0].Member.(string)

	return result_uname, nil
}

func CheckForOpponentAndQueue(user types.User) (string, error) {
	opponent, err := checkForOpponent(user)
	if err != nil {
		return "", err
	}

	if opponent == "" {
		log.Printf("Adding user to queue: %v\n", user)

		RedisClient.ZAdd(context.Background(), "game", redis.Z{Score: float64(user.Rating), Member: user.Username})

		return "", nil
	}

	return opponent, nil
}
