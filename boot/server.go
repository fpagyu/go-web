package boot

import (
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

// AppComponent 组成服务端的应用组件,
// 接口实现者需要提供组件初始化的函数
type AppComponent interface {
	Setup() interface{}
}

type ServerIface interface {
	Serve(addr ...string) error
	RegisteComponent(c AppComponent) error
}

type Components struct {
	*dig.Container
	*viper.Viper
}

func (cs *Components) RegisteComponent(com AppComponent) error {
	t := reflect.TypeOf(com)
	v := reflect.ValueOf(com)

	opts := make([]dig.ProvideOption, 0)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Name == "Name" || field.Name == "Group" {
			vv := v.Field(i)
			if vv.Kind() == reflect.String && vv.String() != "" {
				if field.Name == "Name" {
					opts = append(opts, dig.Name(vv.String()))
				} else {
					opts = append(opts, dig.Group(vv.String()))
				}
			}
		}
	}

	return cs.Container.Provide(com.Setup(), opts...)
}

func ComponentsSetup(coms ...AppComponent) (*Components, error) {
	// 读取项目配置信息
	confPath := []string{
		os.Getenv("CONF_PATH"),
		"config/",
		"../config/",
	}
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	for i := range confPath {
		viper.AddConfigPath(confPath[i])
	}
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cs := &Components{
		Container: dig.New(),
		Viper:     viper.GetViper(),
	}

	_ = cs.Provide(func() *viper.Viper {
		return cs.Viper
	})

	// 注册依赖项
	for i := range coms {
		if err := cs.RegisteComponent(coms[i]); err != nil {
			return nil, err
		}
	}

	return cs, nil
}

type GinApp struct {
	*gin.Engine
	*Components
}

func (s *GinApp) Serve(addrs ...string) (err error) {
	if len(addrs) > 0 {
		err = s.Run(addrs...)
		return
	}

	var addr string
	if s.Viper != nil {
		addr = s.Viper.GetString("addr")
	}
	if addr == "" {
		addr = ":8080"
	}
	err = s.Run(addr)

	return
}

func NewGinApp(components ...AppComponent) (*GinApp, error) {
	coms, err := ComponentsSetup(components...)
	if err != nil {
		return nil, err
	}

	mode := viper.GetString("mode")
	if mode != "" {
		os.Setenv(gin.EnvGinMode, mode)
		gin.SetMode(mode)
	}

	app := &GinApp{
		Engine:     gin.Default(),
		Components: coms,
	}

	return app, nil
}
