package controllers

import (
	v01 "go-web/controllers/v01"
	jwt "go-web/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func RegisteRouter(engine *gin.Engine) {
	engine.Use(jwt.TokenVerify)
	registeRouter_V01(engine)
}

func registeRouter_V01(engine *gin.Engine) {
	r := engine.Group("v01")

	r.GET("hello-world", v01.HelloWorld)
}
