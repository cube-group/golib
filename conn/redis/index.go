package redis

import (
	"github.com/go-redis/redis"
)

type Config struct {
	Address  string
	Password string
	DB       int
	PoolSize int
}

func NewOpen(cfg Config) (*redis.Client, error) {
	conn := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	if err := conn.Ping().Err(); err != nil {
		return nil, err
	}
	return conn, nil
}
