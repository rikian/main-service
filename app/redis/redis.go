package redis

import (
	"log"

	"github.com/go-redis/redis"
)

type RedisService interface {
	GetSessionUser(r *redis.Client, id string) (string, error)
}

type redisImpl struct{}

func InitRedis() RedisService {
	return &redisImpl{}
}

func (s *redisImpl) GetSessionUser(r *redis.Client, id string) (string, error) {
	rDB := r.Get(id)
	if err := rDB.Err(); err != nil {
		log.Print(err.Error())
		return "", err
	}

	result, err := rDB.Result()

	if err != nil {
		log.Print(err.Error())
		return "", err
	}

	return result, nil
}
