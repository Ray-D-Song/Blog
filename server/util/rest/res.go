package rest

import (
	"github.com/gin-gonic/gin"
)

func Res(c *gin.Context, code int, data any) {
	switch code {
	case 200:
		c.JSON(200, gin.H{
			"data": gin.H{
				"code": code,
				"list": data,
			},
		})
	case 404:
		c.JSON(200, gin.H{
			"data": gin.H{
				"code": code,
			},
		})
	case 500:
		c.JSON(200, gin.H{
			"data": gin.H{
				"code": code,
			},
		})
	}
}
