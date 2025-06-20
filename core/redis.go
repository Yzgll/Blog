package core

import (
	"blog/config"
	"blog/global"

	"github.com/go-redis/redis/v8"
)

//redis初始化

var (
	Redis *redis.Client
)

func RedisInit() error {
	RedisDb := redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.Db,
	})
	//ping一下
	_, err := RedisDb.Ping(global.Ctx).Result()
	if err != nil {
		return err
	}
	global.Log.Infof("[redis]连接成功")
	return nil
}
