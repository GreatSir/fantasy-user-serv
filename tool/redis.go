package tool

import (
	"github.com/gomodule/redigo/redis"
	"sync"
	"time"
	"os"
)

var RedisPool *redis.Pool

var one sync.Once

func init()  {
	one.Do(func() {
		host := os.Getenv("REDIS_HOST")
		port := os.Getenv("REDIS_PORT")
		pass := os.Getenv("REDIS_PASS")
		index := os.Getenv("REDIS_DB")
		redisHost := host+":"+port
		RedisPool = &redis.Pool{
			MaxIdle:1,
			MaxActive:10,
			IdleTimeout:180*time.Second,
			Dial: func() (redis.Conn, error) {
				c,err:=redis.Dial("tcp",redisHost)
				if err!=nil{
					return nil,err
				}
				if pass!=""{
					if _,err:=c.Do("AUTH",pass);err!=nil{
						c.Close()
						return nil,err
					}
				}
				if index != "" {
					if _,err := c.Do("SELECT",index);err != nil {
						c.Close()
						return nil,err
					}
				}

				return c,nil
			},
		}
	})
}