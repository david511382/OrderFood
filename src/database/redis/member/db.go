package member

import (
	"github.com/go-redis/redis"
)

type RedisDb struct {
	R *redis.Client
}