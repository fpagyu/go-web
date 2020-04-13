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
	NewFunc() interface{}
}

// ConfigFile 服务端配置文件
type ConfigFile struct {
	Name string   // 配置文件名, 不带扩展名
	Type string   // 配置文件类型
	Path []string // 配置文件所在目录
}

type AppIface interface {
	Serve(addr ...string) error
	RegisteComponent(c AppComponent) error
}

type GinApp struct {
	*gin.Engine
	*dig.Container
	*viper.Viper
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

func (s *GinApp) RegisteComponent(c AppComponent) error {
	t := reflect.TypeOf(c)
	v := reflect.ValueOf(c)

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

	return s.Container.Provide(c.NewFunc(), opts...)
}

func NewGinApp(configFile ConfigFile, components ...AppComponent) (*GinApp, error) {
	viper.SetConfigName(configFile.Name)
	viper.SetConfigType(configFile.Type)
	for i := range configFile.Path {
		viper.AddConfigPath(configFile.Path[i])
	}
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	app := &GinApp{
		Container: dig.New(),
		Viper:     viper.GetViper(),
	}

	_ = app.Container.Provide(func() *viper.Viper {
		return app.Viper
	})

	for i := range components {
		if err := app.RegisteComponent(components[i]); err != nil {
			return nil, err
		}
	}

	mode := viper.GetString("mode")
	if mode != "" {
		os.Setenv(gin.EnvGinMode, mode)
		gin.SetMode(mode)
	}

	app.Engine = gin.Default()

	return app, nil
}
