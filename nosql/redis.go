package nosql

import (
	"github.com/go-redis/redis/v8"
	"sptbackend/config"
)

var Cli *redis.Client

func InitRedis()  {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Cfg.RedisNetwork,
		Password: config.Cfg.RedisPassword, // no password set
		DB:       0,  // use default DB
	})
	Cli = rdb
}




