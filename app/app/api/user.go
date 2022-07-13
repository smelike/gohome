package services

import (
	"github.com/gin-gonic/gin"
)

type User struct {
}

func (that *User) Get(c *gin.Context) {
	var result interface{}

	c.JSON(200, gin.H{
		"code": 200,
		"data": result,
		"msg":  "success",
	})
}
