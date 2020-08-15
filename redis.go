package redis

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	goredis "gopkg.in/redis.v3"
	// goredis "github.com/go-redis/redis"
)

type Config struct {
	Addr     string `default:":6379"`
	Password string // no password set
	DB       int64  `default:"0"` // use default DB
	Prefix   string
}

type RedisClient struct {
	goredis.Client
	prefix string
}

var RDB *RedisClient

func NewRedis(config Config) *RedisClient {
	RDB = &RedisClient{
		Client: *goredis.NewClient(&goredis.Options{
			Addr:     config.Addr,
			Password: config.Password,
			DB:       config.DB,
		}),
		prefix: config.Prefix,
	}

	if pong, err := RDB.Client.Ping().Result(); err != nil {
		panic("redis connect fail:" + err.Error())
	} else {
		fmt.Println("PONG:", pong)
		return nil
	}
	return RDB
}

func (r *RedisClient) FormatKey(args ...interface{}) string {
	return join(r.prefix, join(args...))
}

func join(args ...interface{}) string {
	s := make([]string, len(args))
	for i, v := range args {
		switch v.(type) {
		case string:
			s[i] = v.(string)
		case int64:
			s[i] = strconv.FormatInt(v.(int64), 10)
		case uint64:
			s[i] = strconv.FormatUint(v.(uint64), 10)
		case float64:
			s[i] = strconv.FormatFloat(v.(float64), 'f', 0, 64)
		case bool:
			if v.(bool) {
				s[i] = "1"
			} else {
				s[i] = "0"
			}
		case *big.Int:
			n := v.(*big.Int)
			if n != nil {
				s[i] = n.String()
			} else {
				s[i] = "0"
			}
		default:
			panic("Invalid type specified for conversion")
		}
	}
	return strings.Join(s, ":")
}
