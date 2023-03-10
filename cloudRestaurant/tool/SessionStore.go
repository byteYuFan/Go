package tool

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// InitSession 初始化Session
func InitSession(engine *gin.Engine) {
	config := GetConfig().RedisConfig
	store, err := redis.NewStore(10, "tcp", config.Addr+":"+config.Port, config.Password, []byte("secret"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	engine.Use(sessions.Sessions("mySession", store))
}

// SetSession 设置参数
func SetSession(ctx *gin.Context, key, value any) error {
	session := sessions.Default(ctx)
	if session == nil {
		return nil
	}
	session.Set(key, value)
	return session.Save()
}

// GetSession 获取key
func GetSession(ctx *gin.Context, key any) any {
	session := sessions.Default(ctx)
	return session.Get(key)
}
