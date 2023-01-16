package conn

import (
	"github.com/go-redis/redis/v8"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
)

func GetRedisConn(conf *rkentry.ConfigEntry) *redis.Client {

	addr := conf.GetString("redis_addr")
	pwd := conf.GetString("redis_pwd")
	db := conf.GetInt("redis_db")

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd, // no password set
		DB:       db,  // use default DB
	})

	return rdb
}
