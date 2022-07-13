package services

import (
	"github.com/gin-gonic/gin"
	"github.com/sunquakes/jsonrpc4go"
)

type User struct {
}

func (that *User) Get(c *gin.Context) {
	var result interface{}

	// 调用服务
	cc, _ := jsonrpc4go.NewClient("http", "192.168.0.142", "3233") // the protocol is http
	err := cc.Call("/TestService/Add", map[string]interface{}{
		"a": 1,
		"b": 2,
	}, &result, false)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"code": 400,
			"data": nil,
			"msg":  "error=>" + err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": result,
		"msg":  "success",
	})
}
