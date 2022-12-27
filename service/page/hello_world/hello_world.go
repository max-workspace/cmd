package helloworld

import (
	"github.com/gomodule/redigo/redis"

	"max.workspace.com/cmd/application"
	app_protocol "max.workspace.com/cmd/models/protocol/application"
)

// GetRedisHelloWorld redis使用demo
func (s *Service) GetRedisHelloWorld() (ret string, err error) {
	redisConn, err := application.GetRedisConn(app_protocol.RedisNameTest)
	if err != nil {
		return
	}
	defer redisConn.Close()

	_, err = redisConn.Do("SET", "hello_world", "1")
	if err != nil {
		return
	}
	_, err = redisConn.Do("EXPIRE", "hello_world", 10)
	if err != nil {
		return
	}
	ret, err = redis.String(redisConn.Do("GET", "hello_world"))
	if err != nil {
		return
	}
	return
}
