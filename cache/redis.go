package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/loveyu233/go-utils/client"
	"github.com/loveyu233/go-utils/ctx"
	"github.com/sirupsen/logrus"
)

type RedisConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       int
	Ctx      context.Context
}

var DefaultRedisConfig = RedisConfig{
	Host:     "127.0.0.1",
	Port:     6379,
	Username: "",
	Password: "",
	DB:       0,
}

func MustInitRedisClient(config ...RedisConfig) *redis.Client {
	var (
		err     error
		dConfig RedisConfig
	)

	if len(config) > 0 {
		dConfig = config[0]
	} else {
		dConfig = DefaultRedisConfig
	}

	if dConfig.Ctx == nil {
		dConfig.Ctx = ctx.Timeout()
	}

	client.RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", dConfig.Host, dConfig.Port),
		Username: dConfig.Username,
		Password: dConfig.Password,
		DB:       dConfig.DB,
	})

	if err = client.RedisClient.Ping(dConfig.Ctx).Err(); err != nil {
		logrus.Panicf("ping redis失败: %v", err)
	}

	logrus.Info("redis 连接成功")

	return client.RedisClient
}
