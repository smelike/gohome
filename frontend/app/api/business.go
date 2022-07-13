package services

import (
	"github.com/gin-gonic/gin"
)

type Business struct {
}

func (that *Business) Get(c *gin.Context) {
	// var result interface{}

	// cc, _ :=

	c.JSON(200, gin.H{
		"code": 200,
		"data": true,
		"msg":  "success",
	})
}
