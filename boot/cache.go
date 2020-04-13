package boot

import (
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"github.com/vmihailenco/msgpack"
)

type CacheComponent struct {
	Name  string
	Group string
}

func (c CacheComponent) NewFunc() interface{} {
	return func(conf *viper.Viper) *cache.Codec {
		cfg := conf.GetStringMap("cache")

		ring := redis.NewRing(&redis.RingOptions{
			Addrs:    map[string]string{"default": cfg["addr"].(string)},
			DB:       cfg["db"].(int),
			Password: cfg["passwd"].(string),
		})

		cli := &cache.Codec{
			Redis: ring,
			Marshal: func(v interface{}) ([]byte, error) {
				return msgpack.Marshal(v)
			},
			Unmarshal: func(b []byte, v interface{}) error {
				return msgpack.Unmarshal(b, v)
			},
		}

		return cli
	}
}
