package repositories

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

func StartRedis() error {
	port := os.Getenv("REDIS_PORT")
	if port == "" {
		port = "6379"
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:" + port,
		Password: "",
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

	return rdb.Set(ctx, key, userID, 24*time.Hour).Err()
}

func GetSession(token string) (string, error) {
	key := "session:" + token

	return rdb.Get(ctx, key).Result()
}

func DelSession(token string) error {
	key := "session:" + token

	return rdb.Del(ctx, key).Err()
}
