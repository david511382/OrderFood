package redis

import (
	"orderfood/src/config"
	"orderfood/src/database/common"
	"orderfood/src/database/redis/member"
	"orderfood/src/database/redis/menu"
	"strconv"

	"github.com/go-redis/redis"
)

func NewMemberDb(dbCfg config.DbConfig) (common.IRedisMember, error) {
	rds, err := connect(dbCfg)

	return &member.RedisDb{R: rds}, err
}

func NewMenuDb(dbCfg config.DbConfig) (common.IRedisMenu, error) {
	rds, err := connect(dbCfg)

	return &menu.RedisDb{R: rds}, err
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
