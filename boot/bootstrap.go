package boot

import "os"

var App *GinApp

func init() {
	confPath := os.Getenv("CONF_PATH")

	configFile := ConfigFile{
		Name: "settings",
		Type: "yaml",
		Path: []string{
			confPath,
			"config/",
			"../config/",
		},
	}
	ginApp, err := NewGinApp(
		configFile,
		DBComponent{Name: "db"},
		RedisComponent{Name: "redis"},
		CacheComponent{Name: "cache"},
	)
	if err != nil {
		panic(err)
	}
	App = ginApp
}
