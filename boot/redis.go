package boot

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

type RedisComponent struct {
	Name  string
	Group string
}

func (c RedisComponent) NewFunc() interface{} {
	return func(conf *viper.Viper) (*redis.Client, error) {
		cfg := conf.GetStringMap("redis")

		cli := redis.NewClient(&redis.Options{
			Addr:       cfg["addr"].(string),
			Password:   cfg["passwd"].(string),
			DB:         cfg["db"].(int),
			MaxRetries: cfg["max-retries"].(int),
		})

		_, err := cli.Ping().Result()
		if err != nil {
			println("redis ping err: ", err)
		}

		return cli, nil
	}
}
