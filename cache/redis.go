package cache

import (
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = newPool()
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err)
			}
			return c, err
		},
	}
}
