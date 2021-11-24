package redis_factory

import (
	"context"
	"encoding/json"
	"gebi/app/Http/Serializer"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
)

var Ctx = context.Background()

var RedisClient *redis.Client

func init() {
	gotenv.Load()
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDIS_ADDR"),
		Password:   os.Getenv("REDIS_PW"), // no password set
		DB:         int(db),               // use default DB
		MaxRetries: 1,
	})

	_, err := RedisClient.Ping(Ctx).Result()

	if err != nil {
		defer func() {
			if e := recover(); e != nil {
				log.Printf("recover:%v", e)
			}
		}()
		Serializer.Err(30001, "连接Redis不成功", err)
	}
}

func RedisKey(key string) string {
	pc, _, _, _ := runtime.Caller(1)
	return path.Base(runtime.FuncForPC(pc).Name()) + "_" + key
}

func GetBytes(key string) []byte {
	res, err := RedisClient.Get(Ctx, key).Bytes()
	if err != nil && err != redis.Nil {
		Serializer.Err(30002, "redis读取byte数据失败", err)
	}

	if err == redis.Nil {
		return nil
	}

	return res
}

func SetByJson(key string, value interface{}) {
	data, err := json.Marshal(value)
	if err != nil {
		Serializer.Err(30001, "redis json encode失败", err)
	}

	if RedisClient.Set(Ctx, key, data, time.Hour).Err() != nil {
		Serializer.Err(30001, "设置string缓存失败", err)
	}
}
