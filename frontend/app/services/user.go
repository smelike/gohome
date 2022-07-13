package services

import "github.com/gin-gonic/gin"

type UserService struct {
}

func (that *UserService) Get(c *gin.Context) {

	c.JSON(200, gin.H{
		"code": 200,
		"data": 1,
		"msg":  "success",
	})
}