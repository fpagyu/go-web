package boot

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type DBComponent struct {
	Name  string
	Group string
}

func (c DBComponent) Setup() interface{} {
	return func(conf *viper.Viper) (*gorm.DB, error) {
		cfg := conf.GetStringMapString("db")

		uri := fmt.Sprintf("%s:%s@%s/%s?charset=%s",
			cfg["user"], cfg["passwd"], cfg["addr"], cfg["db-name"], cfg["charset"])

		db, err := gorm.Open(cfg["driver"], uri)
		if err != nil {
			return nil, err
		}

		commDB := db.DB()
		commDB.SetMaxOpenConns(100)
		commDB.SetMaxIdleConns(5)
		commDB.SetConnMaxLifetime(time.Hour)

		if conf.GetString("mode") == "debug" {
			db = db.Debug()
		}

		if err := db.DB().Ping(); err != nil {
			// todo: 标准日志
			println("db init error: ", err)
		}
		return db, nil
	}
}
