package controller

import (
	"server/service"
	"server/util/code"
	"server/util/rest"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBlogList(c *gin.Context) {
	list, err := service.GetBlogList(c)
	if err != nil {
		rest.Res(c, code.Error, nil)
		return
	}
	rest.Res(c, code.OK, list)
}

func GetBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		rest.Res(c, code.Error, nil)
	}
	val, err := service.GetBlog(c, id)
	if err != nil {
		rest.Res(c, code.Error, nil)
	}
	rest.Res(c, code.OK, val)
	_, err = service.AddHit(c, id)
	if err != nil {
	}
}

//func AddBlog(c *gin.Context) {
//	var blog ent.Blog
//	if err := c.ShouldBindJSON(&blog); err != nil {
//		rest.Res(c, code.Error, nil)
//		return
//	}
//}

func AddHit(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		rest.Res(c, code.Error, nil)
	}
	if _, err := service.AddHit(c, id); err != nil {
		rest.Res(c, code.Error, nil)
	}
	rest.Res(c, code.OK, nil)
}
