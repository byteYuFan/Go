package tool

import (
	"context"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type RedisStore struct {
	redisClient *redis.Client
}

var ctx = context.Background()
var RStore RedisStore

func (rs *RedisStore) Set(id string, value string) {
	err := rs.redisClient.Set(ctx, id, value, time.Minute*10).Err()
	if err != nil {
		log.Println(err)
	}
}

func (rs *RedisStore) Get(id string, clear bool) string {
	val, err := rs.redisClient.Get(ctx, id).Result()
	if err != nil {
		log.Println("get.....")
		log.Println(err)
		return ""
	}
	if clear {
		err = rs.redisClient.Del(ctx, id).Err()
		if err != nil {
			fmt.Println("set.....")
			return ""
		}

	}
	return val
}

func InitRedisStore() *RedisStore {
	config := GetConfig().RedisConfig

	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr + ":" + config.Port,
		Password: config.Password,
		DB:       config.Db,
	})
	if client == nil {
		fmt.Println("初始化redis失败")
		return nil
	}
	RStore = RedisStore{
		redisClient: client,
	}
	base64Captcha.SetCustomStore(&RStore)

	return &RStore
}
