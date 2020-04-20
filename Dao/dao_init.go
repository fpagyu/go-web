package dao

import (
	"go-web/boot"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
)

var (
	DB    *gorm.DB
	Redis *redis.Client
	Cache *cache.Codec
)

type daoConns struct {
	dig.In

	DB    *gorm.DB      `name:"db"`
	Redis *redis.Client `name:"redis" optional:"true"`
	Cache *cache.Codec  `name:"cache" optional:"true"`
}

func init() {
	err := boot.App.Invoke(func(v daoConns) {
		DB = v.DB
		Redis = v.Redis
		Cache = v.Cache
	})
	if err != nil {
		panic(err)
	}
}
