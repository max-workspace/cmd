package application

import (
	"time"

	"github.com/gomodule/redigo/redis"

	"max.workspace.com/cmd/models/errors"
	"max.workspace.com/cmd/models/protocol/application"
)

// initRedis redis初始化函数
func initRedis(config application.RedisConfig) (redisPool *redis.Pool, err error) {
	redisPool = &redis.Pool{
		MaxIdle:     config.MaxIdle,
		MaxActive:   config.MaxActive,
		IdleTimeout: time.Second * time.Duration(config.IdleTimeOutS),
		Wait:        config.Wait,
		Dial: func() (c redis.Conn, err error) {
			address := config.Uri
			return redis.Dial("tcp", address,
				redis.DialConnectTimeout(time.Millisecond*time.Duration(config.ConnTimeoutMs)),
				redis.DialReadTimeout(time.Millisecond*time.Duration(config.ReadTimeoutMs)),
				redis.DialWriteTimeout(time.Millisecond*time.Duration(config.WriteTimeoutMs)))
		},
		TestOnBorrow: func(c redis.Conn, tm time.Time) error {
			if time.Since(tm) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	return
}

// GetRedisConn 获取redis连接
func GetRedisConn(redisName string) (redisConn redis.Conn, err error) {
	// 检测全局对象是否初始化
	App := NewApp()
	if App == nil {
		err = errors.ErrorApplicationNotInit
		return
	}

	// 检测redis name是否在初始化的列表里
	redisPool, ok := App.Redis[redisName]
	if !ok {
		err = errors.ErrorRedisNotExist
		return
	}

	// 返回redis的连接
	redisConn = redisPool.Get()
	return
}
