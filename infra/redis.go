package infra

import "github.com/go-redis/redis"

var (
	Redis *redis.Client
)

func InitRedis(addr, pwd string, db int) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
	_, err := Redis.Ping().Result()
	if err != nil {
		panic(err)
	}
}
