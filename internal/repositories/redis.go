package repositories

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

func StartRedis() error {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("REDIS_PORT")
	if port == "" {
		port = "6379"
	}
	password := os.Getenv("REDIS_PASSWORD")

	rdb = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

func SetSession(token string, userID uint) error {
	key := "session:" + token
	timeToLiveStr := os.Getenv("TIME_TO_LIVE")
	timeToLive, err := strconv.Atoi(timeToLiveStr)
	if err != nil {
		return err
	}

	return rdb.Set(ctx, key, userID, time.Duration(timeToLive)*time.Hour).Err()
}

func GetSession(token string) (string, error) {
	key := "session:" + token

	return rdb.Get(ctx, key).Result()
}

func DelSession(token string) error {
	key := "session:" + token

	return rdb.Del(ctx, key).Err()
}

func RefreshSession(token string) error {
	key := "session:" + token
	timeToLiveStr := os.Getenv("TIME_TO_LIVE")
	timeToLive, _ := strconv.Atoi(timeToLiveStr)

	return rdb.Expire(ctx, key, time.Duration(timeToLive)*time.Hour).Err()
}
