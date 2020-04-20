package boot

var App *GinApp

func init() {
	var err error
	App, err = NewGinApp(
		DBComponent{Name: "db"},
		RedisComponent{Name: "redis"},
		CacheComponent{Name: "cache"},
	)
	if err != nil {
		panic(err)
	}
}
