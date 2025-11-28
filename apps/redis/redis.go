package redis

import (
	"context"
	"fmt"
	"tokogue-api/config"

	"github.com/redis/go-redis/v9"
)

// Global Context.
// Library Redis v9 WAJIB butuh context untuk setiap operasinya (biar bisa cancel/timeout).
var Ctx = context.Background()

func ConnectRedis(config config.Config) (*redis.Client, error) {
	// 1. Buat Client Baru
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.Redis.RedisHost, config.Redis.RedisPort),
		Password: config.Redis.RedisPassword,
		DB:       config.Redis.RedisDB, // Default DB biasanya 0
	})

	// 2. Ping untuk Tes Koneksi
	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil

}