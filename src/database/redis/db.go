package redis

import (
	"orderfood/src/config"
	"orderfood/src/database/common"
	"strconv"

	"github.com/go-redis/redis"
)

type redisDb struct {
	r *redis.Client
}

func New(dbCfg config.DbConfig) (common.IRedis, error) {
	rds, err := connect(dbCfg)

	return &redisDb{r: rds}, err
}

func connect(dbCfg config.DbConfig) (*redis.Client, error) {
	db, err := strconv.Atoi(dbCfg.Database)
	if err != nil {
		return nil, err
	}

	// init redis
	redisOpt := &redis.Options{
		Network:  "tcp",
		Addr:     dbCfg.Address,
		Password: dbCfg.Password,
		DB:       db,
		// Dialer func() (net.Conn, error)
		// OnConnect func(*Conn) error
	}

	rds := redis.NewClient(redisOpt)
	err = rds.Ping().Err()
	return rds, err
}
