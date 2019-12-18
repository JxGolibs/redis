package redis

import (
	"github.com/go-redis/redis"
	"fmt"
)

type Config struct {
	Addr     string `default:":6379"`
	Password string // no password set
	DB       int    `default:"0"` // use default DB
}

var RDB *redis.Client

func NewRedis(config Config) *redis.Client {
	RDB := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		 DB:      config.DB,
	})

	if pong, err := RDB.Ping().Result(); err != nil {
		panic("redis connect fail:" + err.Error())
	}else{
	   fmt.Println("PONG:",pong)
	}
	return RDB
}
