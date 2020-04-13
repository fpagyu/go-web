package v01

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	// var d dao.BookDao
	// _, _ = d.FindByID(1)

	c.String(http.StatusOK, "Hello World!!!")
}
