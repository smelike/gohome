package services

import "github.com/gin-gonic/gin"

type BusinessService struct {
}

func (that *BusinessService) Get(c *gin.Context) {

	c.JSON(200, gin.H{
		"code": 200,
		"data": 120000,
		"msg":  "success",
	})
}
