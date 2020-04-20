package jwt

import "github.com/gin-gonic/gin"

func jwtTokenVerify(c *gin.Context) error {
	var token struct {
		Token string `header:"Token"`
	}
	if err := c.BindHeader(&token); err != nil {
		println("bind header err: ", err)
		return err
	}

	// println("token: ", token.Token)
	return nil
}

func TokenVerify(c *gin.Context) {
	_ = jwtTokenVerify(c)
	c.Next()
}
